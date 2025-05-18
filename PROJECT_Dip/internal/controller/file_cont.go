package controller

import (
    "net/http"
    "github.com/Golang/PROJECT_Dip/internal/service"
    "github.com/gin-gonic/gin"
)

type FileController struct {
    fileService *service.FileService
}

func NewFileController(fileService *service.FileService) *FileController {
    return &FileController{fileService: fileService}
}

func (c *FileController) Upload(ctx *gin.Context) {
    userID := ctx.GetUint("userID")
    fileHeader, err := ctx.FormFile("file")
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if fileHeader.Size > 10<<20 { // 10 MB
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "file too large"})
        return
    }

    file, err := c.fileService.UploadFile(userID, fileHeader)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, file)
}

func (c *FileController) ListFiles(ctx *gin.Context) {
    userID := ctx.GetUint("userID")
    files, err := c.fileService.GetUserFiles(userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, files)
}