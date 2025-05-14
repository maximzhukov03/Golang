// Пример кода для создания сервера на go-chi
package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func RouterConfigure(r *chi.Mux) http.Handler{
	r.Use(middleware.Logger)
	r.Get("/1", handleRoute1)
	r.Get("/2", handleRoute2)
	r.Get("/3", handleRoute3)
	return r
}

func main() {
	r := chi.NewRouter()
	RouterConfigure(r)
	http.ListenAndServe(":8080", r)
}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Обработка 1го маршрута"))
}

func handleRoute2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Обработка 2го маршрута"))
}

func handleRoute3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Обработка 3го маршрута"))
}
