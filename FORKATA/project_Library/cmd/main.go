package main

import (
	"database/sql"
	"fmt"
	"golang/project_Library/internal/config"
	handler "golang/project_Library/internal/controller"
	"golang/project_Library/internal/repository"
	"golang/project_Library/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"

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