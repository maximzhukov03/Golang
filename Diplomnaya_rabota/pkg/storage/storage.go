package storage

import (
    "context"
    "database/sql"
    "errors"
    "fmt"
    "mime/multipart"
    "path/filepath"
    "strings"
    "time"

    "github.com/google/uuid"
    "github.com/minio/minio-go/v7"
)

var (
    ErrInvalidFileFormat = errors.New("invalid file format, only png and jpeg allowed")
    ErrFileTooLarge      = errors.New("file size exceeds 10 MB limit")
    ErrUnauthorized      = errors.New("unauthorized user")
)

const maxFileSize = 10 * 1024 * 1024

type Storage interface {
    UploadFile(ctx context.Context, userID int64, fileHeader *multipart.FileHeader) (string, error)
    ListFiles(ctx context.Context, userID int64) ([]FileMeta, error)
}

type storage struct {
    minioClient *minio.Client
    bucketName  string
    db          *sql.DB
}

type FileMeta struct {
    ID        int64
    Name      string
    Size      int64
    Uploaded  time.Time
    URL       string
    UserID    int64
}

func NewStorage(minioClient *minio.Client, bucketName string, db *sql.DB) Storage {
    return &storage{
        minioClient: minioClient,
        bucketName:  bucketName,
        db:          db,
    }
}

func (s *storage) UploadFile(ctx context.Context, userID int64, fileHeader *multipart.FileHeader) (string, error) {
    if userID == 0 {
        return "", ErrUnauthorized
    }
    if fileHeader.Size > maxFileSize {
        return "", ErrFileTooLarge
    }
    ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
    if ext != ".png" && ext != ".jpeg" && ext != ".jpg" {
        return "", ErrInvalidFileFormat
    }
    file, err := fileHeader.Open()
    if err != nil {
        return "", err
    }
    defer file.Close()
    uniqueName := uuid.New().String() + ext
    info, err := s.minioClient.PutObject(ctx, s.bucketName, uniqueName, file, fileHeader.Size, minio.PutObjectOptions{ContentType: getContentType(ext)})
    if err != nil {
        return "", err
    }
    url := fmt.Sprintf("/%s/%s", s.bucketName, uniqueName)
    _, err = s.db.Exec(`
        INSERT INTO files (name, size, uploaded, url, user_id)
        VALUES (?, ?, ?, ?, ?)`,
        uniqueName, info.Size, time.Now().UTC(), url, userID)
    if err != nil {
        return "", err
    }
    return url, nil
}

func (s *storage) ListFiles(ctx context.Context, userID int64) ([]FileMeta, error) {
    rows, err := s.db.Query(`
        SELECT id, name, size, uploaded, url, user_id
        FROM files
        WHERE user_id = ?
        ORDER BY uploaded DESC
    `, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var files []FileMeta
    for rows.Next() {
        var f FileMeta
        err := rows.Scan(&f.ID, &f.Name, &f.Size, &f.Uploaded, &f.URL, &f.UserID)
        if err != nil {
            return nil, err
        }
        files = append(files, f)
    }
    return files, nil
}

func getContentType(ext string) string {
    switch ext {
    case ".png":
        return "image/png"
    case ".jpeg", ".jpg":
        return "image/jpeg"
    default:
        return "application/octet-stream"
    }
}