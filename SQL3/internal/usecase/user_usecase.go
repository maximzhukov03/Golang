package usecase

import (
	"golandg/sql/internal/domain"
)

type UserUseCase interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id int) ([]domain.User, error)
	CreateUser(user domain.User) error
	DeleteUser(user domain.User) error
}

type userUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) GetAllUsers() ([]domain.User, error) {
	return uc.repo.GetAll()
}

func (uc *userUseCase) GetUserByID(id int) ([]domain.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *userUseCase) CreateUser(user domain.User) error {
	return uc.repo.Create(user)
}

func (uc *userUseCase) DeleteUser(user domain.User) error {
	return uc.repo.Delete(user)
}
