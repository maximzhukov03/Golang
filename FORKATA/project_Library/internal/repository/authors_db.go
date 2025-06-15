package database

import (
	"context"
	"database/sql"
	"fmt"
	"golang/project_Library/internal/models"
	"log"
)

type AuthorRepositoryPostgres struct{
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepositoryPostgres{
	return &AuthorRepositoryPostgres{
		db: db,
	}
}

type AuthorRepository interface {
    Create(ctx context.Context ,user models.Author) error
    GetByID(ctx context.Context, id int) (models.Author, error)
    // GetByEmail(ctx context.Context, email string) (models.User, error)
    // GetAll(ctx context.Context) ([]models.User, error)
    GetAllBooks(ctx context.Context, userID int) ([]models.Book, error)
    Delete(ctx context.Context, id int) error
    // Exists(ctx context.Context, id int) (bool, error)
}

func (d *AuthorRepositoryPostgres) Create(ctx context.Context, author models.Author) error{
	query := `INSERT INTO authors (id, name) VALUES ($1, $2)`
	_, err := d.db.ExecContext(ctx, query, author.ID, author.Name)
	if err != nil{
		log.Println("Ошибка в создании Author")
	}
	return err
}

func (d *AuthorRepositoryPostgres) GetByID(ctx context.Context, id int) (models.Author, error){
	query := `SELECT id, name FROM authors WHERE id = $1`
    var a models.Author
    err := d.db.QueryRowContext(ctx, query, id).Scan(&a.ID, &a.Name)
    return a, err
}

func (d *AuthorRepositoryPostgres) GetAllBooks(ctx context.Context, authorID int) ([]models.Book, error) {
    var books []models.Book
    
    query := `
        SELECT b.id, b.title, b.author_id, b.user_id
        FROM books b
        WHERE b.author_id = $1
    `
    
    rows, err := d.db.QueryContext(ctx, query, authorID)
    if err != nil {
        return nil, fmt.Errorf("failed to query author books: %w", err)
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

func (d *AuthorRepositoryPostgres) Delete(ctx context.Context, id int) error{
    query := `DELETE FROM authors WHERE id = $1`
    _, err := d.db.ExecContext(ctx, query, id)
    return err
}	