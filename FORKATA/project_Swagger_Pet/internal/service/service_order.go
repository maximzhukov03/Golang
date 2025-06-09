package service

import (
	"context"
	"errors"
	"golang/project_Swagger_Pet/internal/models"
	"golang/project_Swagger_Pet/internal/repository"
	"time"
)

type OrderService struct {
	repo database.OrderRepository
}

type OrderDTO struct {
	ID       int64
	PetID    int64
	Quantity int
	ShipDate time.Time
}

func NewOrderService(repo database.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, dto OrderDTO) error {
	if dto.PetID == 0 || dto.Quantity <= 0 {
		return errors.New("invalid PetID or Quantity")
	}

	order := models.Order{
		ID:       dto.ID,
		PetID:    dto.PetID,
		Quantity: dto.Quantity,
		ShipDate: dto.ShipDate,
	}

	return s.repo.Create(ctx, order)
}

func (s *OrderService) GetOrderByID(ctx context.Context, id int64) (models.Order, error) {
	if id == 0 {
		return models.Order{}, errors.New("invalid ID")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *OrderService) DeleteOrder(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("invalid ID")
	}
	return s.repo.Delete(ctx, id)
}