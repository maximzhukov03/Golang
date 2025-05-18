package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "net/http"
)

func JWTMiddleware(secret string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        tokenString := ctx.GetHeader("Authorization")
        if tokenString == "" {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            ctx.Set("userID", claims["user_id"])
        } else {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        }
    }
}