package database

import (
	"context"
	"database/sql"
	"log"
)

type User struct{
	ID string
	Name string
	Password string
}

type UserRepositoryPostgres struct{
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *UserRepositoryPostgres{
	return &UserRepositoryPostgres{
		db: db,
	}
}

type UserRepository interface{
	Create(ctx context.Context, id, user_name string, password string) error
	FindUser(ctx context.Context, user_name string) (*User, error)
	Delete(ctx context.Context, user_id string) error
}

func (d *UserRepositoryPostgres) Create(ctx context.Context, id, user_name, password string) error{
	query := `INSERT INTO users (id, name, password) VALUES ($1, $2, $3)`
	_, err := d.db.ExecContext(ctx, query, id, user_name, password)
	if err != nil{
		log.Println("Ошибка в создании БД")
		return err
	}
	return nil
}

func (d *UserRepositoryPostgres) FindUser(ctx context.Context, user_name string) (*User, error){
	user := &User{}
	query := `SELECT id, name, password FROM users WHERE name = $1`
	row := d.db.QueryRowContext(ctx, query, user_name)
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil{
		return nil, err
	}
	return user, nil

}

func (d *UserRepositoryPostgres) Delete(ctx context.Context, user_id string) error{
	query := `DELETE FROM users WHERE id = $1`
	_, err := d.db.ExecContext(ctx, query, user_id)
	if err != nil{
		log.Println("ОШИБКА В УДАЛЕНИИ ПОЛЬЗОВАТЕЛЯ ИЗ БД")
		return err
	}
	return nil
}