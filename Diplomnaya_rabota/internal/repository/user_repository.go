package repository

import (
    "database/sql"
    "errors"
    "golang/myapp/internal/models"
)

// UserRepository описывает методы для работы с пользователями
type UserRepository interface {
    Create(user *models.User) error
    GetByEmail(email string) (*models.User, error)
}

// SQLiteUserRepository реализует UserRepository через SQLite
type SQLiteUserRepository struct {
    db *sql.DB
}

// NewSQLiteUserRepository создаёт новый экземпляр SQLiteUserRepository
func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
    return &SQLiteUserRepository{db: db}
}

// Create добавляет нового пользователя в таблицу users, включая created_at
func (r *SQLiteUserRepository) Create(user *models.User) error {
    query := `INSERT INTO users(email, password_hash, created_at) VALUES(?, ?, ?)`
    res, err := r.db.Exec(query, user.Email, user.PasswordHash, user.CreatedAt)
    if err != nil {
        return err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    user.ID = id
    return nil
}

// GetByEmail возвращает пользователя по email или nil, если не найден
func (r *SQLiteUserRepository) GetByEmail(email string) (*models.User, error) {
    query := `SELECT id, email, password_hash, created_at FROM users WHERE email = ?`
    row := r.db.QueryRow(query, email)
    var user models.User
    if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}