package main

import (
	"context"
	"database/sql"
	"fmt"
	"golang/project_Library/internal/config"
	handler "golang/project_Library/internal/controller"
	"golang/project_Library/internal/models"
	"golang/project_Library/internal/repository"
	"golang/project_Library/internal/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// @title        Library API
// @version      1.0
// @description  Сервис управления пользователями, авторами и книгами.
// @host         localhost:8080
// @BasePath     /
func main() {
	c := config.NewConfig()
	conf := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		c.DB_HOST, c.DB_PORT, c.DB_USER, c.DB_PASSWORD, c.DB_NAME)
	
	db, err := sql.Open("postgres", conf)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Пинг БД: %v", err)
	}

	userRepo := database.NewUserRepository(db)
	authorRepo := database.NewAuthorRepository(db)
	bookRepo := database.NewBooksRepository(db)

	superService := service.NewSuperService(userRepo, authorRepo, bookRepo)

	facade := handler.NewFacade(*superService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/swagger/*", httpSwagger.Handler())
	r.Route("/users", func(r chi.Router) {
		r.Post("/", facade.HandlerCreateUser)
		r.Get("/{id}", facade.HandlerGetUser)
		r.Delete("/{id}", facade.HandlerDelete)

		r.Post("/{idUser}/borrow/{idBook}", facade.HandlerBorrowBook)
		r.Post("/{idUser}/return/{idBook}", facade.HandlerReturnBook)
		r.Get("/{id}/rented", facade.HandlerGetRentedBooks)
	})

	r.Route("/authors", func(r chi.Router) {
		r.Post("/", facade.HandlerCreateAuthor)
		r.Get("/{id}", facade.HandlerGetAuthorByID)
		r.Delete("/{id}", facade.HandlerDeleteAuthor)
		r.Get("/{id}/books", facade.HandlerGetAuthorBooks)
		r.Get("/top", facade.HandlerGetTopAuthors)
	})

	r.Route("/books", func(r chi.Router) {
		r.Post("/", facade.HandlerCreateBook)
		r.Get("/{id}", facade.HandlerGetBookByID)
		r.Delete("/{id}", facade.HandlerDeleteBook)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Сервер запущен на порту :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}


func SeedDatabase(ctx context.Context, userRepo database.UserRepository, authorRepo database.AuthorRepository, bookRepo database.BooksRepository) error {
	gofakeit.Seed(time.Now().UnixNano())

	authors, err := authorRepo.GetAll(ctx)
	if err != nil || len(authors) == 0 {
		fmt.Println("Таблица authors пуста, добавляем авторов")
		for i := 0; i < 10; i++ {
			id := gofakeit.UUID()
			author := models.Author{
				ID:         id,
				Name:       gofakeit.Name(),
				Popularity: gofakeit.Number(0, 100),
			}
			if err := authorRepo.Create(ctx, author); err != nil {
				return fmt.Errorf("не удалось создать автора: %w", err)
			}
		}
	} else {
		fmt.Println("В таблице authors уже есть данные")
	}

	users := []models.User{}
	rows, err := userRepo.GetAll(ctx)
	if err != nil || len(rows) == 0 {
		fmt.Println("Таблица users пуста, добавляем пользователей")
		for i := 0; i < 55; i++ {
			id := gofakeit.UUID()
			user := models.User{
				ID:    id,
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
			}
			if err := userRepo.Create(ctx, user); err != nil {
				return fmt.Errorf("не удалось создать пользователя: %w", err)
			}
			users = append(users, user)
		}
	} else {
		fmt.Println("В таблице users уже есть данные")
		users = rows
	}
	bookCount := 0
	row := bookRepo.Count(ctx)
	if row == 0 {
		fmt.Println("Таблица books пуста, добавляем книги")
		authorsList, err := authorRepo.GetAll(ctx)
		if err != nil || len(authorsList) == 0 {
			return fmt.Errorf("нет авторов для книг")
		}
		for i := 0; i < 100; i++ {
			author := authorsList[gofakeit.Number(0, len(authorsList)-1)]
			userID := "" // часть книг без пользователя (не выдана)
			if gofakeit.Bool() && len(users) > 0 {
				user := users[gofakeit.Number(0, len(users)-1)]
				userID = user.ID
			}
			book := models.Book{
				Title:    gofakeit.BookTitle(),
				AuthorID: author.ID,
				UserID:   userID,
			}
			if err := bookRepo.Create(ctx, book); err != nil {
				return fmt.Errorf("не удалось создать книгу: %w", err)
			}
			bookCount++
		}
	} else {
		fmt.Println("В таблице books уже есть данные")
	}

	fmt.Println("Заполнение базы успешно завершено.")
	return nil
}