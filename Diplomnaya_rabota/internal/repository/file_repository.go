package repository

import (
    "context"
    "database/sql"
    "io"
    "net/url"
    "time"

    "golang/myapp/internal/models"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

// FileRepository описывает методы для работы с метаданными файлов
type FileRepository interface {
    Save(file *models.File) error
    ListByUserID(userID int64) ([]*models.File, error)
}

// StorageClient описывает методы для работы с хранилищем (MinIO)
type StorageClient interface {
    Upload(bucket, objectName string, reader io.Reader, objectSize int64, contentType string) error
    PresignedURL(bucket, objectName string, expiry time.Duration) (string, error)
}

// SQLiteFileRepository реализует FileRepository через SQLite
type SQLiteFileRepository struct {
    db *sql.DB
}

// NewSQLiteFileRepository создает новый экземпляр SQLiteFileRepository
func NewSQLiteFileRepository(db *sql.DB) *SQLiteFileRepository {
    return &SQLiteFileRepository{db: db}
}

// Save сохраняет метаданные файла в таблицу files
func (r *SQLiteFileRepository) Save(file *models.File) error {
    query := `INSERT INTO files(user_id, name, size, bucket, object_name, uploaded_at) VALUES(?, ?, ?, ?, ?, ?)`
    res, err := r.db.Exec(query, file.UserID, file.Name, file.Size, file.Bucket, file.ObjectName, file.UploadedAt)
    if err != nil {
        return err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    file.ID = id
    return nil
}

// ListByUserID возвращает список файлов пользователя
func (r *SQLiteFileRepository) ListByUserID(userID int64) ([]*models.File, error) {
    query := `SELECT id, user_id, name, size, bucket, object_name, uploaded_at FROM files WHERE user_id = ?`
    rows, err := r.db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var files []*models.File
    for rows.Next() {
        var f models.File
        if err := rows.Scan(&f.ID, &f.UserID, &f.Name, &f.Size, &f.Bucket, &f.ObjectName, &f.UploadedAt); err != nil {
            return nil, err
        }
        files = append(files, &f)
    }
    return files, nil
}

// MinIOStorageClient реализует StorageClient через MinIO
type MinIOStorageClient struct {
    client *minio.Client
    bucket string
}

// NewMinIOStorageClient создает новый экземпляр MinIOStorageClient
func NewMinIOStorageClient(endpoint, accessKey, secretKey, bucket string, useSSL bool) (*MinIOStorageClient, error) {
    client, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        return nil, err
    }
    return &MinIOStorageClient{client: client, bucket: bucket}, nil
}

// Upload загружает объект в указанный бакет
func (m *MinIOStorageClient) Upload(bucket, objectName string, reader io.Reader, objectSize int64, contentType string) error {
    _, err := m.client.PutObject(
        context.Background(),
        bucket,
        objectName,
        reader,
        objectSize,
        minio.PutObjectOptions{ContentType: contentType},
    )
    return err
}

// PresignedURL возвращает временную ссылку на объект
func (m *MinIOStorageClient) PresignedURL(bucket, objectName string, expiry time.Duration) (string, error) {
    reqParams := make(url.Values)
    u, err := m.client.PresignedGetObject(context.Background(), bucket, objectName, expiry, reqParams)
    if err != nil {
        return "", err
    }
    return u.String(), nil
}
