package controller

import (
    "net/http"
    "github.com/Golang/PROJECT_Dip/internal/service"
    "github.com/gin-gonic/gin"
)

type AuthController struct {
    authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
    return &AuthController{authService: authService}
}

func (c *AuthController) SignUp(ctx *gin.Context) {
    var request struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=8"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.authService.SignUp(request.Email, request.Password); err != nil {
        ctx.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (c *AuthController) SignIn(ctx *gin.Context) {
    var request struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=8"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := c.authService.SignIn(request.Email, request.Password)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": token})
}