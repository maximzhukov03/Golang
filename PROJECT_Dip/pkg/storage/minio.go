package storage

import (
    "context"
    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
    "io"
)

// NewMinIOClient создает нового клиента MinIO
func NewMinIOClient(endpoint, accessKey, secretKey string) (*minio.Client, error) {
    client, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
        Secure: false,
    })
    if err != nil {
        return nil, err
    }
    
    return client, nil
}

type MinIOStorage struct {
    client     *minio.Client
    bucketName string
}

func NewMinIOStorage(client *minio.Client, bucketName string) *MinIOStorage {
    return &MinIOStorage{client: client, bucketName: bucketName}
}

func (s *MinIOStorage) Upload(ctx context.Context, objectName string, file io.Reader, size int64) (string, error) {
    _, err := s.client.PutObject(
        ctx,
        s.bucketName,
        objectName,
        file,
        size,
        minio.PutObjectOptions{ContentType: "application/octet-stream"},
    )
    if err != nil {
        return "", err
    }

    return s.client.EndpointURL().String() + "/" + s.bucketName + "/" + objectName, nil
}