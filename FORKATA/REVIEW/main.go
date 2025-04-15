package main
import "testing"

type DB interface {
    GetUser(id int) (User, error)
}

type realDB struct{}

func (db *realDB) GetUser(id int) (User, error) {
    // реальная реализация
}

type mockDB struct{}

func (db *mockDB) GetUser(id int) (User, error) {
    return User{ID: id, Name: "Test User"}, nil
}


func TestGetUser(t *testing.T) {
    db := &mockDB{}
    service := NewService(db)
    
    user, err := service.GetUser(1)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    
    if user.Name != "Test User" {
        t.Errorf("unexpected user name: %s", user.Name)
    }
}