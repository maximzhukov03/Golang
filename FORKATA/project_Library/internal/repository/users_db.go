package database

import (
	"context"
	"database/sql"
	"fmt"
	"golang/project_Library/internal/models"
	"log"
)

type UserRepositoryPostgres struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryPostgres{
	return &UserRepositoryPostgres{
		db: db,
	}
}

type UserRepository interface {
    Create(ctx context.Context ,user models.User) error
    GetByID(ctx context.Context, id string) (models.User, error)
    // GetByEmail(ctx context.Context, email string) (models.User, error)
    // GetAll(ctx context.Context) ([]models.User, error)
    GetRentedBooks(ctx context.Context, userID string) ([]models.Book, error)
    Delete(ctx context.Context, id string) error
    GetAll(ctx context.Context) ([]models.User, error)
    // Exists(ctx context.Context, id int) (bool, error)
}

func (d *UserRepositoryPostgres) Create(ctx context.Context, user models.User) error{
	query := `INSERT INTO users (id, name, email) VALUES ($1, $2, $3)`
	_, err := d.db.ExecContext(ctx, query, user.ID, user.Name, user.Email)
	if err != nil{
		log.Println("Ошибка в создании User")
	}
	return err
}

func (d *UserRepositoryPostgres) GetByID(ctx context.Context, id string) (models.User, error){
	query := `SELECT id, name, email FROM users WHERE id = $1`
    var u models.User
    err := d.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Name, &u.Email)
    return u, err
}

func (d *UserRepositoryPostgres) GetRentedBooks(ctx context.Context, userID string) ([]models.Book, error) {
    var books []models.Book
    
    query := `
        SELECT b.id, b.title, b.author_id, b.user_id
        FROM books b
        WHERE b.user_id = $1
    `
    
    rows, err := d.db.QueryContext(ctx, query, userID)
    if err != nil {
        return nil, fmt.Errorf("failed to query rented books: %w", err)
    }
    defer rows.Close()
    
    for rows.Next() {
        var b models.Book
        
        err := rows.Scan(
            &b.ID,
            &b.Title,
            &b.AuthorID,
            &b.UserID,
        )
        if err != nil {
            return nil, fmt.Errorf("failed to scan book: %w", err)
        }
        
        books = append(books, b)
    }
    
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("rows iteration error: %w", err)
    }
    
    return books, nil
}

func (d *UserRepositoryPostgres) Delete(ctx context.Context, id string) error{
    query := `DELETE FROM users WHERE id = $1`
    _, err := d.db.ExecContext(ctx, query, id)
    return err
}	

func (d *UserRepositoryPostgres) GetAll(ctx context.Context) ([]models.User, error) {
    rows, err := d.db.QueryContext(ctx, `SELECT id, name, email FROM users`)
    if err != nil { return nil, err }
    defer rows.Close()
    var result []models.User
    for rows.Next() {
        var u models.User
        if err := rows.Scan(&u.ID, &u.Name, &u.Email); err == nil {
            result = append(result, u)
        }
    }
    return result, nil
}