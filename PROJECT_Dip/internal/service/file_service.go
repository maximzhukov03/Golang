package service

import (
    "context"
    "github.com/Golang/PROJECT_Dip/internal/model"
    "github.com/Golang/PROJECT_Dip/internal/repository"
    "github.com/Golang/PROJECT_Dip/pkg/storage"
    "mime/multipart"
    "time"
)

type FileService struct {
    fileRepo *repository.FileRepository
    storage  storage.Storage
}

func NewFileService(repo *repository.FileRepository, storage storage.Storage) *FileService {
    return &FileService{fileRepo: repo, storage: storage}
}

func (s *FileService) UploadFile(userID uint, fileHeader *multipart.FileHeader) (*model.File, error) {
    file, err := fileHeader.Open()
    if err != nil {
        return nil, err
    }
    defer file.Close()

    objectName := generateObjectName(fileHeader.Filename)
    url, err := s.storage.Upload(context.Background(), objectName, file, fileHeader.Size)
    if err != nil {
        return nil, err
    }

    newFile := &model.File{
        UserID:     userID,
        Filename:   fileHeader.Filename,
        Size:       fileHeader.Size,
        ObjectName: objectName,
        URL:        url,
    }

    if err := s.fileRepo.Create(newFile); err != nil {
        return nil, err
    }

    return newFile, nil
}

func (s *FileService) GetUserFiles(userID uint) ([]model.File, error) {
    return s.fileRepo.GetByUserID(userID)
}

func generateObjectName(filename string) string {
    return time.Now().Format("20060102-150405") + "-" + filename
}