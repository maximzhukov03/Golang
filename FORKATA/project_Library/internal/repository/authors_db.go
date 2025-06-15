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
    Create(ctx context.Context, author models.Author) error
    GetByID(ctx context.Context, id string) (models.Author, error)
    // GetByEmail(ctx context.Context, email string) (models.User, error)
    // GetAll(ctx context.Context) ([]models.User, error)
    GetAllBooks(ctx context.Context, author_id string) ([]models.Book, error)
    Delete(ctx context.Context, id string) error
    UpdateAuthorPopularity(ctx context.Context, authorID string) error
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

func (d *AuthorRepositoryPostgres) GetByID(ctx context.Context, id string) (models.Author, error){
	query := `SELECT id, name FROM authors WHERE id = $1`
    var a models.Author
    err := d.db.QueryRowContext(ctx, query, id).Scan(&a.ID, &a.Name)
    return a, err
}

func (d *AuthorRepositoryPostgres) GetAllBooks(ctx context.Context, authorID string) ([]models.Book, error) {
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

func (d *AuthorRepositoryPostgres) Delete(ctx context.Context, id string) error{
    query := `DELETE FROM authors WHERE id = $1`
    _, err := d.db.ExecContext(ctx, query, id)
    return err
}	

func (d *AuthorRepositoryPostgres) UpdateAuthorPopularity(ctx context.Context, authorID string, newPopularity int) error {
    query := `UPDATE authors SET popularity = $1 WHERE id = $2`
    result, err := d.db.Exec(query, newPopularity, authorID)
    if err != nil {
        return fmt.Errorf("failed to update popularity: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get rows affected: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no author found with id %d", authorID)
    }

    return nil
}

func (d *AuthorRepositoryPostgres) GetTop(ctx context.Context) ([]models.Author, error){
    rows, err := d.db.Query(`SELECT id, name, popularity FROM authors ORDER BY popularity DESC LIMIT 3`)
    if err != nil{
        log.Println("Ошибка на уровне получения топа авторов")
        return nil, err
    }
    defer rows.Close()

    var top []models.Author

    for rows.Next(){
        var author models.Author
        err := rows.Scan(&author.ID, &author.Name, &author.Popularity)
        if err != nil{
            log.Println("ошибка при парсинге строк авторов")
            continue
        }
        top = append(top, author)
    }
    
    return top, nil
}