package storage

import (
    "context"
    "io"
)

type Storage interface {
    Upload(ctx context.Context, objectName string, file io.Reader, size int64) (string, error)
}