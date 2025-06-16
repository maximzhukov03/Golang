package main

import (
    "database/sql"
    "log"
    "time"
    "os"

    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    "golang/myapp/internal/handlers"
    "golang/myapp/internal/middleware"
    "golang/myapp/internal/repository"
    "golang/myapp/internal/service"
    _ "myapp/docs" // swagger generated docs
)

func main() {
    // Загружаем конфиг из env
    sqlitePath := os.Getenv("SQLITE_PATH")
    if sqlitePath == "" {
        sqlitePath = "app.db"
    }
    minioEndpoint := os.Getenv("MINIO_ENDPOINT")
    minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
    minioSecretKey := os.Getenv("MINIO_SECRET_KEY")
    minioBucket := os.Getenv("MINIO_BUCKET")
    useSSL := os.Getenv("MINIO_USE_SSL") == "true"
    jwtSecret := os.Getenv("JWT_SECRET")
    bcryptCost := 12
    urlTTL := time.Minute * 15

    // Подключение к SQLite
    db, err := sql.Open("sqlite3", sqlitePath)
    if err != nil {
        log.Fatalf("failed to open sqlite: %v", err)
    }

    // Инициализируем репозитории
    userRepo := repository.NewSQLiteUserRepository(db)
    fileRepo := repository.NewSQLiteFileRepository(db)

    // Инициализируем клиент MinIO
    minioClient, err := repository.NewMinIOStorageClient(minioEndpoint, minioAccessKey, minioSecretKey, minioBucket, useSSL)
    if err != nil {
        log.Fatalf("failed to init minio client: %v", err)
    }

    // Сервисы
    userService := service.NewUserService(userRepo, bcryptCost)
    fileService := service.NewFileService(fileRepo, minioClient, minioBucket, urlTTL)

    // Хендлеры
    authHandler := handler.NewAuthHandler(userService)
    fileHandler := handler.NewFileHandler(fileService)

    // Gin
    r := gin.Default()

    // Swagger UI
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Публичные эндпоинты
    r.POST("/api/register", authHandler.Register)
    r.POST("/api/login", authHandler.Login)

    // Защищённые
    api := r.Group("/api")
    api.Use(middleware.JWTAuth())
    api.POST("/files", fileHandler.UploadFile)
    api.GET("/files", fileHandler.ListFiles)

    // Запуск
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run("0.0.0.0:" + port)
}
