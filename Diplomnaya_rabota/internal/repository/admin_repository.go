package repository

import (
    "database/sql"
    "errors"
    "golang/myapp/internal/models"
)

// AdminRepository описывает CRUD операции для модели User, доступные администратору
type AdminRepository interface {
    List() ([]*models.User, error)
    GetByID(id int64) (*models.User, error)
    Update(user *models.User) error
    Delete(id int64) (int64, error)
}

// SQLiteAdminRepository реализует AdminRepository с помощью SQLite
type SQLiteAdminRepository struct {
    db *sql.DB
}

// NewSQLiteAdminRepository создаёт новый SQLiteAdminRepository
func NewSQLiteAdminRepository(db *sql.DB) *SQLiteAdminRepository {
    return &SQLiteAdminRepository{db: db}
}

// List возвращает всех пользователей
func (r *SQLiteAdminRepository) List() ([]*models.User, error) {
    rows, err := r.db.Query(`SELECT id, email, password_hash, created_at, role FROM users`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*models.User
    for rows.Next() {
        var u models.User
        if err := rows.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.Role); err != nil {
            return nil, err
        }
        u.PasswordHash = ""
        users = append(users, &u)
    }
    return users, nil
}

// GetByID возвращает пользователя по его ID или nil, если не найден
func (r *SQLiteAdminRepository) GetByID(id int64) (*models.User, error) {
    row := r.db.QueryRow(`SELECT id, email, password_hash, created_at, role FROM users WHERE id = ?`, id)
    var u models.User
    if err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.Role); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    u.PasswordHash = ""
    return &u, nil
}

// Update изменяет поля email и роль пользователя
func (r *SQLiteAdminRepository) Update(user *models.User) error {
    res, err := r.db.Exec(`UPDATE users SET email = ?, role = ? WHERE id = ?`, user.Email, user.Role, user.ID)
    if err != nil {
        return err
    }
    affected, err := res.RowsAffected()
    if err != nil {
        return err
    }
    if affected == 0 {
        return errors.New("no rows updated")
    }
    return nil
}

// Delete удаляет пользователя по ID, возвращает число удалённых строк
func (r *SQLiteAdminRepository) Delete(id int64) (int64, error) {
    res, err := r.db.Exec(`DELETE FROM users WHERE id = ?`, id)
    if err != nil {
        return 0, err
    }
    return res.RowsAffected()
}