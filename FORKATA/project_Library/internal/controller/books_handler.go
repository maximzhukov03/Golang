package handler

import (
	"encoding/json"
	"fmt"
	"golang/project_Library/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// HandlerCreateBook создаёт книгу.
// @Summary     Создать книгу
// @Tags        Books
// @Accept      json
// @Produce     json
// @Param       book  body      models.Book  true  "Данные книги"
// @Success     200   {object}  handler.Response{data=models.Book}
// @Failure     400   {object}  handler.Response
// @Failure     500   {object}  handler.Response
// @Router      /books [post]
func (h *Facade) HandlerCreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		h.Responder.ErrorBadRequest(w, fmt.Errorf("invalid request body: %w", err))
		return
	}

	if book.Title == "" || book.AuthorID == "" {
		h.Responder.ErrorBadRequest(w, fmt.Errorf("title and author_id are required"))
		return
	}

	err = h.FacadeService.CreateBook(r.Context(), book)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Data:    book,
	})
}

// HandlerGetBookByID возвращает книгу по ID.
// @Summary     Получить книгу
// @Tags        Books
// @Produce     json
// @Param       id   path      string  true  "ID книги"
// @Success     200  {object}  handler.Response{data=models.Book}
// @Failure     400  {object}  handler.Response
// @Failure     500  {object}  handler.Response
// @Router      /books/{id} [get]
func (h *Facade) HandlerGetBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := h.FacadeService.GetBookByID(r.Context(), id)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Data:    book,
	})
}

// HandlerDeleteBook удаляет книгу.
// @Summary     Удалить книгу
// @Tags        Books
// @Produce     json
// @Param       id   path      string  true  "ID книги"
// @Success     200  {object}  handler.Response
// @Failure     400  {object}  handler.Response
// @Failure     500  {object}  handler.Response
// @Router      /books/{id} [delete]
func (h *Facade) HandlerDeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.FacadeService.Delete(r.Context(), id)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Message: "Book deleted successfully",
	})
}