package service

import (
    "errors"

    "golang/myapp/internal/models"
    "golang/myapp/internal/repository"
)

// AdminService описывает бизнес-логику для операций администратора с пользователями
// Включает в себя операции списка, получения, обновления, удаления и повышения роли пользователя

type AdminService interface {
    ListUsers() ([]*models.User, error)
    GetUserByID(id int64) (*models.User, error)
    UpdateUser(id int64, email, role string) (*models.User, error)
    DeleteUser(id int64) (bool, error)
    PromoteUser(id int64) (*models.User, error)
}

// adminService реализует AdminService
type adminService struct {
    repo repository.AdminRepository
}

// NewAdminService создаёт новый экземпляр AdminService
func NewAdminService(repo repository.AdminRepository) AdminService {
    return &adminService{repo: repo}
}

// ListUsers возвращает список всех пользователей
func (s *adminService) ListUsers() ([]*models.User, error) {
    return s.repo.List()
}

// GetUserByID возвращает пользователя по его ID
func (s *adminService) GetUserByID(id int64) (*models.User, error) {
    return s.repo.GetByID(id)
}

// UpdateUser обновляет email и роль пользователя и возвращает обновлённую сущность
func (s *adminService) UpdateUser(id int64, email, role string) (*models.User, error) {
    user, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, nil
    }
    user.Email = email
    user.Role = role
    if err := s.repo.Update(user); err != nil {
        return nil, err
    }
    return user, nil
}

// DeleteUser удаляет пользователя по ID, возвращает true, если удаление произошло
func (s *adminService) DeleteUser(id int64) (bool, error) {
    rows, err := s.repo.Delete(id)
    if err != nil {
        return false, err
    }
    return rows > 0, nil
}

// PromoteUser повышает пользователя до роли "admin"
func (s *adminService) PromoteUser(id int64) (*models.User, error) {
    user, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, nil
    }
    if user.Role == "admin" {
        return user, errors.New("user is already admin")
    }
    user.Role = "admin"
    if err := s.repo.Update(user); err != nil {
        return nil, err
    }
    return user, nil
}