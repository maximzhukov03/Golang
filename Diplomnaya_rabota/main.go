package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"

    "golang/diplom/internal/db"
    "golang/diplom/internal/user"
    "golang/diplom/pkg/auth"
    "golang/diplom/pkg/storage"
)

const (
    jwtSecret     = "supersecret"
    sqlitePath    = "app.db"
    minioEndpoint = "localhost:9000"
    minioBucket   = "uploads"
    minioKey      = "minioadmin"
    minioSecret   = "minioadmin"
)

func main() {
    database, err := db.InitDB(sqlitePath)
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    minioClient, err := minio.New(minioEndpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(minioKey, minioSecret, ""),
        Secure: false,
    })
    if err != nil {
        log.Fatalf("Failed to connect to MinIO: %v", err)
    }

    ctx := context.Background()
    err = minioClient.MakeBucket(ctx, minioBucket, minio.MakeBucketOptions{})
    if err != nil {
        exists, errBucketExists := minioClient.BucketExists(ctx, minioBucket)
        if errBucketExists != nil || !exists {
            log.Fatalf("Could not create or verify bucket: %v", err)
        }
    }

    userService := user.NewService(database, jwtSecret)
    fileStorage := storage.NewStorage(minioClient, minioBucket, database)

    r := gin.Default()

    r.POST("/register", func(c *gin.Context) {
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
            return
        }

        token, err := userService.RegisterUser(req.Email, req.Password, jwtSecret)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": token})
    })

    r.POST("/login", func(c *gin.Context) {
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
            return
        }

        token, err := userService.AuthenticateUser(req.Email, req.Password, jwtSecret)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": token})
    })

    authMiddleware := func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            return
        }

        userID, err := auth.ParseJWT(token, jwtSecret)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        c.Set("userID", userID)
        c.Next()
    }

    authGroup := r.Group("/api", authMiddleware)

    authGroup.POST("/upload", func(c *gin.Context) {
        userID := c.GetInt64("userID")

        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "missing file"})
            return
        }

        url, err := fileStorage.UploadFile(c, userID, file)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"url": url})
    })

    authGroup.GET("/files", func(c *gin.Context) {
        userID := c.GetInt64("userID")

        files, err := fileStorage.ListFiles(c, userID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch files"})
            return
        }

        c.JSON(http.StatusOK, files)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    srv := &http.Server{
        Addr:         ":" + port,
        Handler:      r,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    log.Printf("Server running on :%s", port)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("server error: %v", err)
    }
}