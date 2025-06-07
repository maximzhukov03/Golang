package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang/project_API/internal/config"
	"golang/project_API/internal/controller"
	"golang/project_API/internal/repository"
	"golang/project_API/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	c := config.NewConfig()
	conf := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.DB_HOST, c.DB_PORT, c.DB_USER, c.DB_PASSWORD, c.DB_NAME)
	db, err := sql.Open("postgres", conf)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Пинг БД: %v", err)
	}

	repo := database.NewDb(db)
	userService := service.NewUserService(repo)
	responder := handler.NewResponder()
	userHandler := handler.NewHandleUser(userService, responder)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.HandlerListUsers)
		r.Post("/", userHandler.HandlerCreateUser)
		r.Get("/{id}", userHandler.HandlerGetId)
		r.Put("/", userHandler.HandlerUpdateUser)
		r.Delete("/{id}", userHandler.HandlerDelete)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Сервер запущен на порту :%s", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}