package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"golandg/sql/internal/delivery/http/handlers"
	"golandg/sql/internal/repository/postgres"
	"golandg/sql/internal/usecase"
)

func main() {
	// Подключение к базе данных
	connect := "host=127.0.0.1 port=5432 user=postgres dbname=users_log sslmode=disable password=goLANG"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")

	// Инициализация слоев
	userRepo := postgres.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	handler := handlers.NewHandler(userUseCase)

	// Настройка маршрутов
	mux := http.NewServeMux()
	mux.HandleFunc("/user", handler.HandleUser)

	// Запуск сервера
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server starting on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
