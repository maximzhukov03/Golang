package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang/project_Swagger_Pet/internal/config"
	"golang/project_Swagger_Pet/internal/controller"
	"golang/project_Swagger_Pet/internal/repository"
	"golang/project_Swagger_Pet/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "golang/project_Swagger_Pet/docs"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

// @title           Swagger Petstore API
// @version         1.0
// @description     This is a sample server Petstore server.
// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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

	userRepo := database.NewUserDb(db)
	petRepo := database.NewPetDb(db)
	orderRepo := database.NewOrderDb(db)

	userService := service.NewUserService(userRepo)
	petService := service.NewPetService(petRepo)
	orderService := service.NewOrderService(orderRepo)

	responder := handler.NewResponder()

	userHandler := handler.NewHandleUser(userService, responder)
	petHandler := handler.NewHandlePet(petService, responder)
	orderHandler := handler.NewHandleOrder(orderService, responder)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/swagger/*", httpSwagger.Handler())
	r.Route("/users", func(r chi.Router) {
		r.Get("/login", userHandler.HandlerLogin)
		r.Post("/", userHandler.HandlerCreateUser)
		r.Get("/{username}", userHandler.HandlerGetUser)
	})
	r.Group(func(r chi.Router) {
		r.Use(userHandler.JWTAuthMiddleware)
		r.Put("/users", userHandler.HandlerUpdateUser)
		r.Delete("/users/{username}", userHandler.HandlerDelete)
		
		r.Route("/pets", func(r chi.Router) {
			r.Post("/", petHandler.HandlerCreatePet)
			r.Get("/", petHandler.HandlerFindByStatus)
			r.Get("/{id}", petHandler.HandlerGetPet)
			r.Put("/", petHandler.HandlerUpdatePet)
			r.Delete("/{id}", petHandler.HandlerDeletePet)
		})
		
		r.Route("/orders", func(r chi.Router) {
			r.Post("/", orderHandler.HandlerCreateOrder)
			r.Get("/{id}", orderHandler.HandlerGetOrder)
			r.Delete("/{id}", orderHandler.HandlerDeleteOrder)
		})
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