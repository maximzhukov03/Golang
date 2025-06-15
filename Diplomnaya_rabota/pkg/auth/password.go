package auth

import (
    "errors"
    "unicode"

    "golang.org/x/crypto/bcrypt"
)

var (
    ErrWeakPassword = errors.New("password must be at least 8 characters and contain letters and digits")
)

func HashPassword(password string) (string, error) {
    const cost = 12
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    return string(hashed), err
}

func ComparePassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidatePassword(password string) error {
    if len(password) < 8 {
        return ErrWeakPassword
    }
    var hasLetter, hasDigit bool
    for _, c := range password {
        if unicode.IsLetter(c) {
            hasLetter = true
        } else if unicode.IsDigit(c) {
            hasDigit = true
        }
    }
    if !hasLetter || !hasDigit {
        return ErrWeakPassword
    }
    return nil
}