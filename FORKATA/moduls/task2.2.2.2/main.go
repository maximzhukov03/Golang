package main

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func NewUser(id int, options ...UserOption) *User{
	o := &User{
		ID: id,
	}
	
	
	for _, option := range options{
		option(o)
	}

	return o
}

func WithUsername(str string) UserOption{
	return func(u *User) {
		u.Username = str
	}
}

func WithEmail(str string) UserOption{
	return func(u *User) {
		u.Email = str
	}
}

func WithRole(str string) UserOption{
	return func(u *User) {
		u.Role = str
	}
}

func main() {
	user := NewUser(1, WithUsername("testuser"), WithEmail("testuser@example.com"), WithRole("admin"))
	fmt.Printf("User: %+v\n", user)
}