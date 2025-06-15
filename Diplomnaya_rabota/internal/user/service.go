package user

import (
    "database/sql"
    "errors"
    "strings"
    "time"
    "net/mail"

    "golang.org/x/crypto/bcrypt"
    "golang/diplom/pkg/auth"
)

type Service interface {
    Register(email, password string) error
    Authenticate(email, password string) (string, error)
}

type service struct {
    db         *sql.DB
    jwtSecret  string
}

func NewService(db *sql.DB, jwtSecret string) Service {
    return &service{db: db, jwtSecret: jwtSecret}
}

var (
    ErrInvalidEmail       = errors.New("invalid email format")
    ErrWeakPassword       = errors.New("password must be at least 8 characters and contain letters and digits")
    ErrUserExists         = errors.New("user already exists")
    ErrInvalidCredentials = errors.New("invalid email or password")
)

func (s *service) Register(email, password string) error {
    email = strings.TrimSpace(email)
    if err := validateEmail(email); err != nil {
        return err
    }
    if err := validatePassword(password); err != nil {
        return err
    }

    if exists, _ := s.emailExists(email); exists {
        return ErrUserExists
    }

    hashed, err := hashPassword(password)
    if err != nil {
        return err
    }

    return s.insertUser(email, hashed)
}

func (s *service) Authenticate(email, password string) (string, error) {
    user, err := s.getUserByEmail(email)
    if err != nil {
        return "", ErrInvalidCredentials
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", ErrInvalidCredentials
    }

    token, err := auth.GenerateJWT(user.ID, s.jwtSecret)
    if err != nil {
        return "", err
    }

    return token, nil
}

func validateEmail(email string) error {
    _, err := mail.ParseAddress(email)
    if err != nil || len(email) > 320 {
        return ErrInvalidEmail
    }
    return nil
}

func validatePassword(password string) error {
    if len(password) < 8 {
        return ErrWeakPassword
    }
    var hasLetter, hasDigit bool
    for _, c := range password {
        switch {
        case 'a' <= c && c <= 'z', 'A' <= c && c <= 'Z':
            hasLetter = true
        case '0' <= c && c <= '9':
            hasDigit = true
        }
    }
    if !hasLetter || !hasDigit {
        return ErrWeakPassword
    }
    return nil
}

func hashPassword(password string) (string, error) {
    const cost = 12
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    return string(hashed), err
}

func (s *service) emailExists(email string) (bool, error) {
    var exists int
    err := s.db.QueryRow("SELECT 1 FROM users WHERE email = ? LIMIT 1", email).Scan(&exists)
    if err == sql.ErrNoRows {
        return false, nil
    }
    if err != nil {
        return false, err
    }
    return true, nil
}

func (s *service) insertUser(email, hashedPassword string) error {
    _, err := s.db.Exec(`
        INSERT INTO users (email, password, created_at)
        VALUES (?, ?, ?)
    `, email, hashedPassword, time.Now().UTC())
    return err
}

func (s *service) getUserByEmail(email string) (*User, error) {
    user := &User{}
    row := s.db.QueryRow(`
        SELECT id, email, password, created_at
        FROM users
        WHERE email = ?
    `, email)

    err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
    if err != nil {
        return nil, err
    }

    return user, nil
}