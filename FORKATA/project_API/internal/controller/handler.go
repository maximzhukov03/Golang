package handler

import (
	"encoding/json"
	"fmt"
	"golang/project_API/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type Respond struct{}

func NewResponder() Responder {
	return &Respond{}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println("responder json encode error:", err)
	}
}

func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	log.Println("http response bad request status code:", err)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		log.Println("response writer error on write:", err)
	}
}

func (r *Respond) ErrorInternal(w http.ResponseWriter, err error) {
	log.Println("http response internal server error:", err)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: err.Error(),
	}); err != nil {
		log.Println("response writer error on write:", err)
	}
}

type HandleUser struct {
	handler   *service.UserService
	responder Responder
}

func NewHandleUser(service *service.UserService, responder Responder) *HandleUser {
	return &HandleUser{
		handler:   service,
		responder: responder,
	}
}

// @Summary      Get User by ID
// @Description  get user
// @Accept       json
// @Produce      json
// @Param        id  path string true  "User id"
// @Success      200  {object}  Response
// @Failure      500  {string}  Response
// @Router       /users/{id} [get]
func (h *HandleUser) HandlerGetId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.handler.GetUser(r.Context(), id)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    user,
	})
}

// @Summary      Create User
// @Description  create
// @Accept       json
// @Produce      json
// @Param        user  body service.UserStruct true  "User data"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users [post]
func (h *HandleUser) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var user service.UserStruct

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		log.Println("Ошибка в обработке тела запроса")
		h.responder.ErrorBadRequest(w, fmt.Errorf("query parameter is required"))
		return
	}

	err = h.handler.CreateUser(r.Context(), user)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return	
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    user,
	})
}

// @Summary      Update User data Base
// @Description  update
// @Accept       json
// @Produce      json
// @Param        user  body service.UserStruct true  "User data"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users [put]
func (h *HandleUser) HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	var user service.UserStruct

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		log.Println("Ошибка в обработке тела запроса")
		h.responder.ErrorBadRequest(w, fmt.Errorf("query parameter is required"))
		return
	}

	err = h.handler.UpdateUser(r.Context(), user)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return	
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    user,
	})
}

// @Summary      Update User data Base
// @Description  update
// @Accept       json
// @Produce      json
// @Param        id  path string true  "User id"
// @Success      200  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users/{id} [delete]
func (h *HandleUser) HandlerDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.handler.DeleteUser(r.Context(), id)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
	})
}


// @Summary      List of users
// @Description  get users
// @Accept       json
// @Produce      json
// @Param        limit  query int false  "Лимит" default(10)
// @Param        offset query int false  "Смещение" default(0)
// @Success      200  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users [get]
func (h *HandleUser) HandlerListUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")

	limit := 10
	offset := 0

	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	conditions := service.ConditionsStruct{
		Limit:  limit,
		Offset: offset,
	}


	users, err := h.handler.List(r.Context(), conditions)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data: map[string]interface{}{
			"total": len(users),
			"users": users,
		},
	})
}