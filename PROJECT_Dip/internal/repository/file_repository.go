package repository

import (
    "gorm.io/gorm"
    "github.com/Golang/PROJECT_Dip/internal/model"
)

type FileRepository struct {
    db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
    return &FileRepository{db: db}
}

func (r *FileRepository) Create(file *model.File) error {
    return r.db.Create(file).Error
}

func (r *FileRepository) GetByUserID(userID uint) ([]model.File, error) {
    var files []model.File
    err := r.db.Where("user_id = ?", userID).Find(&files).Error
    return files, err
}