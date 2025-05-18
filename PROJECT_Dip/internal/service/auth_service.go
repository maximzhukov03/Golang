package service

import (
	"time"
	"github.com/Golang/PROJECT_Dip/internal/model"
	"github.com/Golang/PROJECT_Dip/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    userRepo *repository.UserRepository
    jwtSecret string
}

func NewAuthService(repo *repository.UserRepository, secret string) *AuthService {
    return &AuthService{userRepo: repo, jwtSecret: secret}
}

func (s *AuthService) SignUp(email, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := &model.User{
        Email:        email,
        PasswordHash: string(hashedPassword),
    }

    return s.userRepo.Create(user)
}

func (s *AuthService) SignIn(email, password string) (string, error) {
    user, err := s.userRepo.GetByEmail(email)
    if err != nil {
        return "", err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return "", err
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
    })

    return token.SignedString([]byte(s.jwtSecret))
}