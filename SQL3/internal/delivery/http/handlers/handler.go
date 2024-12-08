package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"golandg/sql/internal/domain"
	"golandg/sql/internal/usecase"
)

type Handler struct {
	userUseCase usecase.UserUseCase
}

func NewHandler(userUseCase usecase.UserUseCase) *Handler {
	return &Handler{
		userUseCase: userUseCase,
	}
}

func (h *Handler) HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getUsers(w, r)
	case http.MethodPost:
		h.createUser(w, r)
	case http.MethodDelete:
		h.deleteUser(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUseCase.GetAllUsers()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(users)
	if err != nil {
		log.Printf("Error marshaling users: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(reqBytes, &user); err != nil {
		log.Printf("Error unmarshaling user: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.userUseCase.CreateUser(user); err != nil {
		log.Printf("Error creating user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(reqBytes, &user); err != nil {
		log.Printf("Error unmarshaling user: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.userUseCase.DeleteUser(user); err != nil {
		log.Printf("Error deleting user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
