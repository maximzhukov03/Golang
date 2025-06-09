package service

import (
	"context"
	"errors"
	"golang/project_Swagger_Pet/internal/models"
	"golang/project_Swagger_Pet/internal/repository"
)

type PetService struct {
	repo database.PetRepository
}

type PetDTO struct {
	ID     int64
	Name   string
	Status string
}

func NewPetService(repo database.PetRepository) *PetService {
	return &PetService{repo: repo}
}

func (s *PetService) CreatePet(ctx context.Context, dto PetDTO) error {
	if dto.Name == "" || dto.Status == "" {
		return errors.New("name or status is empty")
	}

	pet := models.Pet{
		ID:     dto.ID,
		Name:   dto.Name,
		Status: dto.Status,
	}

	return s.repo.Create(ctx, pet)
}

func (s *PetService) UpdatePet(ctx context.Context, dto PetDTO) error {
	if dto.ID == 0 || dto.Name == "" || dto.Status == "" {
		return errors.New("id, name or status is empty")
	}

	pet := models.Pet{
		ID:     dto.ID,
		Name:   dto.Name,
		Status: dto.Status,
	}

	return s.repo.Update(ctx, pet)
}

func (s *PetService) DeletePet(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	return s.repo.Delete(ctx, id)
}

func (s *PetService) GetPetByID(ctx context.Context, id int64) (models.Pet, error) {
	if id == 0 {
		return models.Pet{}, errors.New("id is empty")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *PetService) GetPetsByStatus(ctx context.Context, status string) ([]models.Pet, error) {
	if status == "" {
		return nil, errors.New("status is empty")
	}
	return s.repo.FindByStatus(ctx, status)
}

