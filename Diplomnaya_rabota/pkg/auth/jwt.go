package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
    ErrInvalidToken = errors.New("invalid token")
)

func GenerateJWT(userID int64, secret string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })

    return token.SignedString([]byte(secret))
}

func ParseJWT(tokenStr string, secret string) (int64, error) {
    token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, ErrInvalidToken
        }
        return []byte(secret), nil
    })

    if err != nil || !token.Valid {
        return 0, ErrInvalidToken
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, ErrInvalidToken
    }
    expFloat, ok := claims["exp"].(float64)
    if !ok || int64(expFloat) < time.Now().Unix() {
        return 0, ErrInvalidToken
    }
    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return 0, ErrInvalidToken
    }

    return int64(userIDFloat), nil
}