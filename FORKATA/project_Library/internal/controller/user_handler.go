package handler

import (
	"encoding/json"
	"fmt"
	"golang/project_Library/internal/models"
	"golang/project_Library/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Facade struct{
	FacadeService *service.SuperService
	Responder Responder
}

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

func NewFacade(service service.SuperService) *Facade {
	return &Facade{
		FacadeService: &service,
		Responder:     NewResponder(),
	}
}

// HandlerGetUser возвращает пользователя по ID.
// @Summary     Получить пользователя
// @Tags        Users
// @Produce     json
// @Param       id   path      string  true  "ID пользователя"
// @Success     200  {object}  handler.Response{data=models.User}
// @Failure     400  {object}  handler.Response
// @Failure     500  {object}  handler.Response
// @Router      /users/{id} [get]
func (h *Facade) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.FacadeService.GetUserByID(r.Context(), id)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Data:    user,
	})
}

// HandlerCreateUser создаёт нового пользователя.
// @Summary     Создать пользователя
// @Tags        Users
// @Accept      json
// @Produce     json
// @Param       user  body      models.User  true  "Данные пользователя"
// @Success     200   {object}  handler.Response{data=models.User}
// @Failure     400   {object}  handler.Response
// @Failure     500   {object}  handler.Response
// @Router      /users [post]
func (h *Facade) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		log.Println("Ошибка в обработке тела запроса")
		h.Responder.ErrorBadRequest(w, fmt.Errorf("query parameter is required"))
		return
	}

	err = h.FacadeService.CreateUser(r.Context(), user)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return	
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Data:    user,
	})
}	

// HandlerDelete удаляет пользователя по ID.
// @Summary     Удалить пользователя
// @Tags        Users
// @Produce     json
// @Param       id   path      string  true  "ID пользователя"
// @Success     200  {object}  handler.Response
// @Failure     400  {object}  handler.Response
// @Failure     500  {object}  handler.Response
// @Router      /users/{id} [delete]
func (h *Facade) HandlerDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.FacadeService.DeleteUser(r.Context(), id)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
	})
}

// HandlerBorrowBook выдаёт книгу пользователю.
// @Summary     Выдать книгу
// @Tags        Users
// @Produce     json
// @Param       idUser  path      string  true  "ID пользователя"
// @Param       idBook  path      string  true  "ID книги"
// @Success     200     {object}  handler.Response
// @Failure     400     {object}  handler.Response
// @Failure     500     {object}  handler.Response
// @Router      /users/{idUser}/borrow/{idBook} [post]
func (h *Facade) HandlerBorrowBook(w http.ResponseWriter, r *http.Request) {
	idUser := chi.URLParam(r, "idUser")
	idBook := chi.URLParam(r, "idBook")

	err := h.FacadeService.BorrowBook(r.Context(), idUser, idBook)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
	})
}

// HandlerReturnBook возвращает книгу от пользователя.
// @Summary     Вернуть книгу
// @Tags        Users
// @Produce     json
// @Param       idUser  path      string  true  "ID пользователя"
// @Param       idBook  path      string  true  "ID книги"
// @Success     200     {object}  handler.Response
// @Failure     400     {object}  handler.Response
// @Failure     500     {object}  handler.Response
// @Router      /users/{idUser}/return/{idBook} [post]
func (h *Facade) HandlerReturnBook(w http.ResponseWriter, r *http.Request) {
	idUser := chi.URLParam(r, "idUser")
	idBook := chi.URLParam(r, "idBook")

	err := h.FacadeService.ReturnBook(r.Context(), idUser, idBook)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
	})
}

// HandlerGetRentedBooks возвращает выданные книги пользователя.
// @Summary     Список выданных книг
// @Tags        Users
// @Produce     json
// @Param       id   path      string  true  "ID пользователя"
// @Success     200  {object}  handler.Response{data=[]models.Book}
// @Failure     400  {object}  handler.Response
// @Failure     500  {object}  handler.Response
// @Router      /users/{id}/rented [get]
func (h *Facade) HandlerGetRentedBooks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.FacadeService.GetRentedBooks(r.Context(), id)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Data:    user,
	})
}