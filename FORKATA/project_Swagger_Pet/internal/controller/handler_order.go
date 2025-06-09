package handler

import (
	"encoding/json"
	"fmt"
	"golang/project_Swagger_Pet/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type HandleOrder struct {
	handler   *service.OrderService
	responder Responder
}

func NewHandleOrder(service *service.OrderService, responder Responder) *HandleOrder {
	return &HandleOrder{
		handler:   service,
		responder: responder,
	}
}

// @Summary      Get Order by ID
// @Description  get order
// @Accept       json
// @Produce      json
// @Param        id  path int true  "Order ID"
// @Success      200  {object}  Response
// @Failure      500  {object}  Response
// @Router       /orders/{id} [get]
func (h *HandleOrder) HandlerGetOrder(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.responder.ErrorBadRequest(w, fmt.Errorf("invalid order id"))
		return
	}

	order, err := h.handler.GetOrderByID(r.Context(), id)
	if err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    order,
	})
}

// @Summary      Create Order
// @Description  create new order
// @Accept       json
// @Produce      json
// @Param        order  body service.OrderDTO true  "Order data"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /orders [post]
func (h *HandleOrder) HandlerCreateOrder(w http.ResponseWriter, r *http.Request) {
	var order service.OrderDTO

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		log.Println("Ошибка при декодировании тела запроса:", err)
		h.responder.ErrorBadRequest(w, fmt.Errorf("invalid request body"))
		return
	}

	if err := h.handler.CreateOrder(r.Context(), order); err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Data:    order,
	})
}

// @Summary      Delete Order
// @Description  delete order by ID
// @Accept       json
// @Produce      json
// @Param        id  path int true  "Order ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /orders/{id} [delete]
func (h *HandleOrder) HandlerDeleteOrder(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.responder.ErrorBadRequest(w, fmt.Errorf("invalid order id"))
		return
	}

	if err := h.handler.DeleteOrder(r.Context(), id); err != nil {
		h.responder.ErrorInternal(w, err)
		return
	}

	h.responder.OutputJSON(w, Response{
		Success: true,
		Message: "Order deleted successfully",
	})
}