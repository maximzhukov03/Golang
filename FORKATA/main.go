package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"reflect"
	"strings"
)

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type Tabler interface {
	TableName() string
}

func (u User) TableName() string{
	return "User"
}



type SQLiteGenerator struct{}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL() string
	CreateInsertSQL() string
}

func (s *SQLiteGenerator) CreateTableSQL(u *User) string{
	t := reflect.TypeOf(*u)
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", u.TableName())
	queryHelp := ""
	for i := 0; i < t.NumField(); i++{
		field := t.Field(i)
		dbfield := field.Tag.Get("db_field")
		dbtype := t.Field(i).Tag.Get("db_type")

		if dbfield != "" && dbtype != ""{
			if i > 0{
				queryHelp += ","
			}	

			queryHelp += fmt.Sprintf("%s %s", dbfield, dbtype)
		}
	}

	query += queryHelp + ");"

	return query

}

func (s *SQLiteGenerator) CreateInsertSQL(t Tabler) string {
	var sb strings.Builder
	val := reflect.ValueOf(t).Elem()
	typ := val.Type()

	var fields []string
	var values []string

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		dbfield := field.Tag.Get("db_field")
		dbvalue := val.Field(i).Interface()

		fields = append(fields, dbfield)
		values = append(values, fmt.Sprintf("'%v'", dbvalue))
	}

	sb.WriteString(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", t.TableName(), strings.Join(fields, ", "), strings.Join(values, ", ")))
	return sb.String()
}
type GoFakeitGenerator struct{}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() string
}

func (f *GoFakeitGenerator) GenerateFakeUser() User{
	u := User{
		FirstName: gofakeit.FirstName(),
		LastName: gofakeit.LastName(),
		Email: gofakeit.Email(),
	}
	return u
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