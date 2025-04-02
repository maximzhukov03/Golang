package main

import (
	"fmt"
	// "github.com/brianvoe/gofakeit/v6"
	// "reflect"
	// "strings"
)

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type SQLiteGenerator struct{}
type GoFakeitGenerator struct{}

type Tabler interface {
	TableName() string
}

func (u User) TableName() string{
	panic("implement me")
}


type SQLGenerator interface {
	CreateTableSQL(Tabler) string
	CreateInsertSQL(Tabler) string
}

type FakeDataGenerator interface {
	GenerateFakeUser() User
}

func (s *SQLiteGenerator) CreateTableSQL(t Tabler) string{
	panic("implement me")
}

func (s *SQLiteGenerator) CreateInsertSQL(t Tabler) string{
	panic("implement me")
}

func (s *GoFakeitGenerator) GenerateFakeUser() User{
	panic("implement me")
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
}