package service

import (
	"context"

	"fmt"
	"golang/project_API/internal/repository"
	"log"
)

type UserService struct {
    repo database.UserRepository
}

type UserStruct struct{
	Name string
	Email string
}

type ConditionsStruct struct{
	Limit int
	Offset int
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

	user := database.User{
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

	user := database.User{
		Name: userH.Name,
		Email: userH.Email,
	}

	return service.repo.Update(ctx, user)
}	


func (service *UserService) DeleteUser(ctx context.Context, id string) error{
	if id == ""{
		log.Println("id are empty")
		return nil
	}

	return service.repo.Delete(ctx, id)
}	

func (service *UserService) GetUser(ctx context.Context, id string) (database.User, error){
	if id == ""{
		log.Println("id are empty")
		return database.User{}, nil
	}

	return service.repo.GetByID(ctx, id)
}	

func (service *UserService) List(ctx context.Context, cH ConditionsStruct) ([]database.User, error){
	c := database.Conditions{
		Limit: cH.Limit,
		Offset: cH.Offset,
	}

	return service.repo.List(ctx, c)
}