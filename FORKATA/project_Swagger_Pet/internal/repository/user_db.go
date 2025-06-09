package database

import (
	"context"
	"database/sql"
	"log"
	"golang/project_Swagger_Pet/internal/models"
)

type UserRepositoryPostgres struct {
	db *sql.DB
}

func NewUserDb(db *sql.DB) *UserRepositoryPostgres{
	return &UserRepositoryPostgres{
		db: db,
	}
}

type UserRepository interface {
    Create(ctx context.Context, user models.User) error
    GetByUsername(ctx context.Context, username string) (models.User, error)
    Update(ctx context.Context, user models.User) error
    Delete(ctx context.Context, username string) error
    GetByCredentials(ctx context.Context, username, password string) (models.User, error)
}


func (d *UserRepositoryPostgres) Create(ctx context.Context, user models.User) error{
	query := `INSERT INTO users (id, name, email) VALUES ($1, $2, $3)`
	_, err := d.db.ExecContext(ctx, query, user.ID, user.Name, user.Email)
	if err != nil{
		log.Println("Ошибка в создании User")
		return err
	}
	return nil
}

func (d *UserRepositoryPostgres) GetByUsername(ctx context.Context, username string) (models.User, error) {
	query := `SELECT id, name, email FROM users WHERE name = $1 AND is_deleted = FALSE`
	var u models.User
	err := d.db.QueryRowContext(ctx, query, username).Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}

func (d *UserRepositoryPostgres) Delete(ctx context.Context, username string) error {
	query := `UPDATE users SET is_deleted = TRUE WHERE name = $1`
	_, err := d.db.ExecContext(ctx, query, username)
	return err
}

func (d *UserRepositoryPostgres) Update(ctx context.Context, user models.User) error{
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3;`
	_, err := d.db.ExecContext(ctx, query, user.Name, user.Email, user.ID)
	if err != nil{
		log.Println("Ошибка в Update User")
		return err
	}
	return nil
}

func (d *UserRepositoryPostgres) GetByCredentials(ctx context.Context, username, password string) (models.User, error) {
    query := `SELECT id, name, email FROM users WHERE is_deleted = FALSE AND name = $1 AND password = $2`

	var u models.User
	err := d.db.QueryRowContext(ctx, query, username, password).Scan(&u.ID, &u.Name, &u.Email)
if err != nil {
	log.Printf("Ошибка при получении пользователя по учетным данным: %v", err)
}
	return u, err
}