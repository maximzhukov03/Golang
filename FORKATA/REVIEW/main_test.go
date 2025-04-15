package service

import (
    "testing"
    "errors"
)

type MockUserRepository struct {
    users map[int]*User
    shouldFail bool
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
    if m.shouldFail {
        return nil, errors.New("database error")
    }
    return m.users[id], nil
}

func TestGetUserName_Success(t *testing.T) {
    mockRepo := &MockUserRepository{
        users: map[int]*User{
            1: {ID: 1, Name: "Alice"},
        },
    }
    service := &UserService{repo: mockRepo}

    name, err := service.GetUserName(1)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if name != "Alice" {
        t.Errorf("Expected 'Alice', got '%s'", name)
    }
}

func TestGetUserName_Error(t *testing.T) {
    mockRepo := &MockUserRepository{shouldFail: true}
    service := &UserService{repo: mockRepo}

    _, err := service.GetUserName(1)
    if err == nil {
        t.Error("Expected error, got nil")
    }
}