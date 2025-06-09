package handler

import (
	"encoding/json"
	"fmt"
	"golang/project_Swagger_Pet/internal/models"
	"golang/project_Swagger_Pet/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type HandlePet struct {
	handler   *service.PetService
	responder Responder
}

func NewHandlePet(service *service.PetService, responder Responder) *HandlePet {
	return &HandlePet{
		handler:   service,
		responder: responder,
	}
}

// @Summary Create a new pet
// @Accept json
// @Produce json
// @Param pet body models.Pet true "Pet object"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /pets [post]
func (h *HandlePet) HandlerCreatePet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		h.responder.ErrorBadRequest(w, fmt.Errorf("invalid request body"))
		return
	}

	pets := service.PetDTO{
		ID:     pet.ID,
		Name:   pet.Name,
		Status: pet.Status,
	}

	if err := h.handler.CreatePet(r.Context(), pets); err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    pets,
	})
}

// @Summary Get pet by ID
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /pets/{id} [get]
func (h *HandlePet) HandlerGetPet(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.responder.ErrorBadRequest(w, fmt.Errorf("invalid ID"))
		return
	}

	pet, err := h.handler.GetPetByID(r.Context(), id)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    pet,
	})
}

// @Summary Update pet
// @Accept json
// @Produce json
// @Param pet body models.Pet true "Pet object"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /pets [put]
func (h *HandlePet) HandlerUpdatePet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		h.responder.ErrorBadRequest(w, fmt.Errorf("invalid request body"))
		return
	}
	pets := service.PetDTO{
		ID:     pet.ID,
		Name:   pet.Name,
		Status: pet.Status,
	}

	if err := h.handler.UpdatePet(r.Context(), pets); err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    pets,
	})
}

// @Summary Delete pet
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /pets/{id} [delete]
func (h *HandlePet) HandlerDeletePet(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.responder.ErrorBadRequest(w, fmt.Errorf("invalid ID"))
		return
	}

	if err := h.handler.DeletePet(r.Context(), id); err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{Success: true})
}

// @Summary Find pets by status
// @Produce json
// @Param status query string true "Pet status"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /pets [get]
func (h *HandlePet) HandlerFindByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		h.responder.ErrorBadRequest(w, fmt.Errorf("status query param is required"))
		return
	}

	pets, err := h.handler.GetPetsByStatus(r.Context(), status)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    pets,
	})
}