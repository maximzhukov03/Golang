package handler

import (
	"encoding/json"
	"fmt"
	"golang/project_Library/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// HandlerCreateAuthor создаёт автора.
// @Summary     Создать автора
// @Tags        Authors
// @Accept      json
// @Produce     json
// @Param       author  body      models.Author  true  "Данные автора"
// @Success     200     {object}  handler.Response{data=models.Author}
// @Failure     400     {object}  handler.Response
// @Failure     500     {object}  handler.Response
// @Router      /authors [post]
func (h *Facade) HandlerCreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		h.Responder.ErrorBadRequest(w, fmt.Errorf("invalid request body: %w", err))
		return
	}

	if author.Name == "" {
		h.Responder.ErrorBadRequest(w, fmt.Errorf("author name is required"))
		return
	}

	err = h.FacadeService.CreateAuthor(r.Context(), author)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Data:    author,
	})
}

// HandlerGetAuthorByID возвращает автора по ID.
// @Summary     Получить автора
// @Tags        Authors
// @Produce     json
// @Param       id   path      string  true  "ID автора"
// @Success     200  {object}  handler.Response{data=models.Author}
// @Failure     400  {object}  handler.Response
// @Failure     500  {object}  handler.Response
// @Router      /authors/{id} [get]
func (h *Facade) HandlerGetAuthorByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	author, err := h.FacadeService.GetAuthorByID(r.Context(), id)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Data:    author,
	})
}

// HandlerDeleteAuthor удаляет автора.
// @Summary     Удалить автора
// @Tags        Authors
// @Produce     json
// @Param       id   path      string  true  "ID автора"
// @Success     200  {object}  handler.Response
// @Failure     400  {object}  handler.Response
// @Failure     500  {object}  handler.Response
// @Router      /authors/{id} [delete]
func (h *Facade) HandlerDeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.FacadeService.DeleteAuthor(r.Context(), id)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Message: "Author deleted successfully",
	})
}


// HandlerGetAuthorBooks возвращает книги автора.
// @Summary     Книги автора
// @Tags        Authors
// @Produce     json
// @Param       id   path      string  true  "ID автора"
// @Success     200  {object}  handler.Response{data=[]models.Book}
// @Failure     400  {object}  handler.Response
// @Failure     500  {object}  handler.Response
// @Router      /authors/{id}/books [get]
func (h *Facade) HandlerGetAuthorBooks(w http.ResponseWriter, r *http.Request) {
	authorID := chi.URLParam(r, "id")

	books, err := h.FacadeService.GetAllAuthorsBooks(r.Context(), authorID)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Data:    books,
	})
}

// HandlerGetTopAuthors возвращает топ авторов.
// @Summary     Топ авторов
// @Tags        Authors
// @Produce     json
// @Success     200  {object}  handler.Response{data=[]models.Author}
// @Failure     500  {object}  handler.Response
// @Router      /authors/top [get]
func (h *Facade) HandlerGetTopAuthors(w http.ResponseWriter, r *http.Request) {
	topAuthors, err := h.FacadeService.GetTopAuthors(r.Context())
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}

	h.Responder.OutputJSON(w, Response{
		Success: true,
		Data:    topAuthors,
	})
}