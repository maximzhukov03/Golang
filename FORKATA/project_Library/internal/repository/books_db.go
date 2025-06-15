package database

import (
	"context"
	"database/sql"
	"golang/project_Library/internal/models"
	"log"
)

type BooksRepositoryPostgres struct{
	db *sql.DB
}

func NewBooksRepository(db *sql.DB) *BooksRepositoryPostgres{
	return &BooksRepositoryPostgres{
		db: db,
	}
}

type BooksRepository interface {
    Create(ctx context.Context , book models.Book) error
    GetByID(ctx context.Context, id string) (models.Book, error)
    Delete(ctx context.Context, id string) error
	Update(ctx context.Context, book models.Book) error
}

func (d *BooksRepositoryPostgres) Create(ctx context.Context, book models.Book) error{
	query := `INSERT INTO books (id, title, author_id, user_id) VALUES ($1, $2, $3, $4)`
	_, err := d.db.ExecContext(ctx, query, book.ID, book.Title, book.AuthorID, book.UserID)
	if err != nil{
		log.Println("Ошибка в создании Book")
	}
	return err
}

func (d *BooksRepositoryPostgres) GetByID(ctx context.Context, id string) (models.Book, error){
	query := `SELECT id, title, author_id, user_id FROM books WHERE id = $1`
    var b models.Book
    err := d.db.QueryRowContext(ctx, query, id).Scan(&b.ID, &b.Title, &b.AuthorID, &b.UserID)
    return b, err
}

func (d *BooksRepositoryPostgres) Delete(ctx context.Context, id string) error{
    query := `DELETE FROM books WHERE id = $1`
    _, err := d.db.ExecContext(ctx, query, id)
    return err
}	

func (r *BooksRepositoryPostgres) Update(ctx context.Context, book models.Book) error {
	query := `
		UPDATE books
		SET title = $1, author_id = $2, user_id = $3
		WHERE id = $4
	`
	_, err := r.db.ExecContext(ctx, query, book.Title, book.AuthorID, book.UserID, book.ID)
	if err != nil {
		log.Println("Ошибка при обновлении книги:", err)
	}
	return err
}