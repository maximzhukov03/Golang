package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
    JwtSecret = []byte("hu2iohfiudshf09h72qoifuhsdjfh")
    JwtTTL    = time.Hour * 24
)

type Claims struct {
    UserID string `json:"user_id"`
    jwt.RegisteredClaims
}

func GenerateToken(userID string) (string, error) {
    now := time.Now()
    claims := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            IssuedAt:  jwt.NewNumericDate(now),
            ExpiresAt: jwt.NewNumericDate(now.Add(JwtTTL)),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JwtSecret)
}

func ParseToken(tokenString string) (*Claims, error) {
    tok, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
        return JwtSecret, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := tok.Claims.(*Claims); ok && tok.Valid {
        return claims, nil
    }
    return nil, jwt.ErrTokenInvalidClaims
}

type ctxKey string

const UserIDKey ctxKey = "user_id"

func RequireAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        hdr := r.Header.Get("Authorization")
        if !strings.HasPrefix(hdr, "Bearer ") {
            http.Error(w, "Authorization header missing or invalid", http.StatusUnauthorized)
            return
        }
        tokenString := strings.TrimPrefix(hdr, "Bearer ")
        claims, err := ParseToken(tokenString)
        if err != nil {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }
        ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}