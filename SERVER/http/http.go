package http

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "http_server/docs"
)

type Storage interface {
	Get(key string) (*string, error)
	Put(key string, value string) error
	Post(key string, value string) error
	Delete(key string) error
}

type Server struct {
	storage Storage
}

func newServer(storage Storage) *Server {
	return &Server{storage: storage}
}

// @Summary Get a value
// @Description Get a value by key
// @Param key query string true "Key"
// @Success 200 {string} string "Value"
// @Failure 400 {string} string "Missing Key"
// @Router /object [get]
func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}
	value, err := s.storage.Get(key)
	if err != nil || value == nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	_, _ = fmt.Fprintf(w, *value)
}

func (s *Server) putHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	key, okKey := data["key"]
	value, okValue := data["value"]

	if !okValue || !okKey {
		http.Error(w, "Missing key or value", http.StatusBadRequest)
	}

	if err := s.storage.Put(key, value); err != nil {
		http.Error(w, "Failed to store value", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	key, okKey := data["key"]
	value, okValue := data["value"]

	if !okValue || !okKey {
		http.Error(w, "Missing key or value", http.StatusBadRequest)
	}

	if err := s.storage.Post(key, value); err != nil {
		http.Error(w, "Failed to store value", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	if err := s.storage.Delete(key); err != nil {
		http.Error(w, "Failed to delete key", http.StatusInternalServerError)
		return
	}
}

func CreateAndRunServer(storage Storage, addr string) error {
	server := newServer(storage)

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/object", func(r chi.Router) {
		r.Get("/", server.getHandler)
		r.Post("/", server.postHandler)
		r.Put("/", server.putHandler)
		r.Delete("/", server.deleteHandler)
	})

	httpServer := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return httpServer.ListenAndServe()
}
