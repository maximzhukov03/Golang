package main

import (
    "database/sql"
    "log"
    "os"
    "path/filepath"
    "time"

    "github.com/gin-gonic/gin"
    _ "github.com/glebarez/sqlite" // pure Go SQLite driver, works with CGO_ENABLED=0
    "github.com/pressly/goose/v3"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    "golang/myapp/internal/handlers"
    "golang/myapp/internal/middleware"
    "golang/myapp/internal/repository"
    "golang/myapp/internal/service"
    _ "golang/myapp/docs" // swagger generated docs
)

// @title           MyApp API
// @version         1.0
// @description     API для управления пользователями и файлами
// @securityDefinitions.apikey  bearerAuth
// @in                          header
// @name                        Authorization
func main() {
    // Load config from env
    sqlitePath := os.Getenv("SQLITE_PATH")
    if sqlitePath == "" {
        sqlitePath = "app.db"
    }
    minioEndpoint := os.Getenv("MINIO_ENDPOINT")
    if minioEndpoint == "" {
        minioEndpoint = "localhost:9000"
    }
    minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
    minioSecretKey := os.Getenv("MINIO_SECRET_KEY")
    minioBucket := os.Getenv("MINIO_BUCKET")
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        jwtSecret = "dsfajij4098jasdkf"
    }
    useSSL := os.Getenv("MINIO_USE_SSL") == "true"
    bcryptCost := 12
    urlTTL := time.Minute * 15

    // Open SQLite database using pure Go driver
    db, err := sql.Open("sqlite", sqlitePath)
    if err != nil {
        log.Fatalf("failed to open sqlite: %v", err)
    }

    // Apply migrations
    migrationsDir := filepath.Join(".", "migrations")
    if err := goose.SetDialect("sqlite3"); err != nil {
        log.Fatalf("goose.SetDialect: %v", err)
    }
    if err := goose.Up(db, migrationsDir); err != nil {
        log.Fatalf("goose.Up: %v", err)
    }

    // Initialize repositories
    userRepo := repository.NewSQLiteUserRepository(db)
    fileRepo := repository.NewSQLiteFileRepository(db)
    adminRepo := repository.NewSQLiteAdminRepository(db)

    // Initialize MinIO client
    minioClient, err := repository.NewMinIOStorageClient(minioEndpoint, minioAccessKey, minioSecretKey, minioBucket, useSSL)
    if err != nil {
        log.Fatalf("failed to init minio client: %v", err)
    }

    // Initialize services
    userService := service.NewUserService(userRepo, bcryptCost)
    fileService := service.NewFileService(fileRepo, minioClient, minioBucket, urlTTL)
    adminService := service.NewAdminService(adminRepo)

    // Initialize handlers
    authHandler := handler.NewAuthHandler(userService)
    fileHandler := handler.NewFileHandler(fileService)
    adminHandler := handler.NewAdminHandler(adminService)

    // Set up Gin router
    r := gin.Default()

    // Swagger UI
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Public endpoints
    r.POST("/api/register", authHandler.Register)
    r.POST("/api/login", authHandler.Login)

    // Protected endpoints
    api := r.Group("/api")
    api.Use(middleware.JWTAuth(jwtSecret))
    {
        api.POST("/files", fileHandler.UploadFile)
        api.GET("/files", fileHandler.ListFiles)
    }

    // Admin endpoints (require admin role)
    admin := r.Group("/api/admin")
    admin.Use(middleware.JWTAuth(jwtSecret))
    // TODO: add middleware.RequireAdmin(adminService) to enforce role
    {
        admin.GET("/users",             adminHandler.ListUsers)
        admin.GET("/users/:id",         adminHandler.GetUser)
        admin.PUT("/users/:id",         adminHandler.UpdateUser)
        admin.DELETE("/users/:id",      adminHandler.DeleteUser)
        admin.POST("/users/:id/promote", adminHandler.PromoteUser)
    }

    // Run server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run("0.0.0.0:" + port)
}
