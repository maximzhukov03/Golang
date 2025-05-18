package model

import "gorm.io/gorm"

type File struct {
    gorm.Model
    UserID     uint
    Filename   string
    Size       int64
    ObjectName string // Имя файла в MinIO
    URL        string
}