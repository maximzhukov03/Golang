package service

import (
	"context"

	"fmt"
	"golang/project_Swagger_Pet/internal/repository"
	"golang/project_Swagger_Pet/internal/models"
	"log"
)

type UserService struct {
    repo database.UserRepository
}

type UserStruct struct{
	Name string
	Email string
}

func NewUserService(repo database.UserRepository) *UserService{
	return &UserService{
		repo: repo,
	}
}

func (service *UserService) CreateUser(ctx context.Context, userH UserStruct) error{
	if userH.Name == "" || userH.Email == ""{
		log.Println("Пустые поля имя или почты")
		return fmt.Errorf("name or email are empty")
	}

	user := models.User{
		Name: userH.Name,
		Email: userH.Email,
	}


	return service.repo.Create(ctx, user)
}	

func (service *UserService) UpdateUser(ctx context.Context, userH UserStruct) error{
	if userH.Name == "" || userH.Email == ""{
		log.Println("name or email are empty")
		return nil
	}

	user := models.User{
		Name: userH.Name,
		Email: userH.Email,
	}

	return service.repo.Update(ctx, user)
}	


func (service *UserService) DeleteUser(ctx context.Context, username string) error{
	if username == ""{
		log.Println("id are empty")
		return nil
	}

	return service.repo.Delete(ctx, username)
}	

func (service *UserService) GetUser(ctx context.Context, username string) (models.User, error){
	if username == ""{
		log.Println("username are empty")
		return models.User{}, nil
	}

	return service.repo.GetByUsername(ctx, username)
}	

func (service *UserService) GetByCredentials(ctx context.Context, username, password string) (models.User, error){
	if username == "" || password == ""{
		log.Println("username are empty")
		return models.User{}, nil
	}

	return service.repo.GetByCredentials(ctx, username, password)
}

