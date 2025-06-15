package service

import (
	"context"
	"fmt"
	"golang/project_Library/internal/models"
	"golang/project_Library/internal/repository"
	"log"
)

type SuperService struct{
	UserService *database.UserRepositoryPostgres
	AuthorService *database.AuthorRepositoryPostgres
	BookService *database.BooksRepositoryPostgres
}


func NewSuperService(userDB *database.UserRepositoryPostgres, authorDB *database.AuthorRepositoryPostgres, booksDB *database.BooksRepositoryPostgres) *SuperService{
	return &SuperService{
		UserService: userDB,
		AuthorService: authorDB,
		BookService: booksDB,
	}
}

func (service *SuperService) CreateUser(ctx context.Context, user models.User) error{
	if user.Name == "" || user.Email == ""{
		log.Println("Проблема с созданием пользователя в сервисе")
		return fmt.Errorf("Invalid name or email")
	}
	return service.UserService.Create(ctx, user)
}

func (service *SuperService) GetUserByID(ctx context.Context, id string) (models.User, error){
	return service.UserService.GetByID(ctx, id)
}

func (service *SuperService) GetRentedBooks(ctx context.Context, userID string) ([]models.Book, error){
	return service.UserService.GetRentedBooks(ctx, userID)
}

func (service *SuperService) DeleteUser(ctx context.Context, userID string) error{
	return service.UserService.Delete(ctx, userID)
}

func (service *SuperService) BorrowBook(ctx context.Context, userID, bookID string) error{
	var user models.User
	var book models.Book
	var err error
	user, err = service.UserService.GetByID(ctx, userID)
	if err != nil{
		log.Println("ОШибка в проверке существования пользователя")
		return err
	}
	book, err = service.BookService.GetByID(ctx, bookID)
	if err != nil{
		log.Println("ОШибка в проверке существования книги")
		return err
	}

	if book.UserID != "0"{
		return fmt.Errorf("Book is already borrowed")
	}

	book.UserID = user.ID

	err = service.BookService.Update(ctx, book)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении книги: %w", err)
	}

	result, err := service.AuthorService.GetByID(ctx, book.AuthorID)
	if err != nil{
		return err
	}

	err = service.AuthorService.UpdateAuthorPopularity(ctx, book.AuthorID, result.Popularity+1)
	if err != nil{
		return err
	}
	return nil
}

func (service *SuperService) ReturnBook(ctx context.Context, userID, bookID string) error{
	var book models.Book
	var err error
	_, err = service.UserService.GetByID(ctx, userID)
	if err != nil{
		log.Println("Ошибка в проверке существования пользователя")
		return err
	}
	book, err = service.BookService.GetByID(ctx, bookID)
	if err != nil{
		log.Println("Ошибка в проверке существования книги")
		return err
	}

	if book.UserID != userID{
		log.Println("Ошибка в проверке существования пользователя")
		return fmt.Errorf("user not rented book")
	}

	if book.UserID == "0"{
		return fmt.Errorf("Book is not already borrowed")
	}

	book.UserID = "0"

	err = service.BookService.Update(ctx, book)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении книги: %w", err)
	}

	return nil
}