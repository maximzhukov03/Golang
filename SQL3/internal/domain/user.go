package domain

import (
	"encoding/json"
)

// User представляет сущность пользователя
type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

// UserRepository интерфейс для работы с хранилищем пользователей
type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int) ([]User, error)
	Create(user User) error
	Delete(user User) error
}
