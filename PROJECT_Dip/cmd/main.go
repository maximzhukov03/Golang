package main

import (
    "github.com/Golang/PROJECT_Dip/internal/config"
    "github.com/Golang/PROJECT_Dip/internal/controller"
    "github.com/Golang/PROJECT_Dip/internal/middleware"
    "github.com/Golang/PROJECT_Dip/internal/model"
    "github.com/Golang/PROJECT_Dip/internal/repository"
    "github.com/Golang/PROJECT_Dip/internal/service"
    "github.com/Golang/PROJECT_Dip/pkg/storage"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

func main() {
    cfg := config.Load()
    
    // Инициализация БД
    db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&model.User{}, &model.File{})

    // Инициализация MinIO
    minioClient, err := storage.NewMinIOClient(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey)
    if err != nil {
        log.Fatalf("Failed to initialize MinIO client: %v", err)
    }

    // Инициализация репозиториев
    userRepo := repository.NewUserRepository(db)
    fileRepo := repository.NewFileRepository(db)

    // Инициализация сервисов
    authService := service.NewAuthService(userRepo, cfg.JWTSecret)
    fileService := service.NewFileService(fileRepo, storage.NewMinIOStorage(minioClient, "uploads"))

    // Инициализация контроллеров
    authController := controller.NewAuthController(authService)
    fileController := controller.NewFileController(fileService)

    // Настройка маршрутов
    router := gin.Default()

    // Public routes
    router.POST("/sign-up", authController.SignUp)
    router.POST("/sign-in", authController.SignIn)

    // Protected routes
    auth := router.Group("/")
    auth.Use(middleware.JWTMiddleware(cfg.JWTSecret))
    {
        auth.POST("/upload", fileController.Upload)
        auth.GET("/files", fileController.ListFiles)
    }

    router.Run(":8080")
}