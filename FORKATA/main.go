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
	CreateTableSQL(Tabler) string
	CreateInsertSQL(Tabler) string
}

func (s *SQLiteGenerator) CreateTableSQL(tab Tabler) string{
	t := reflect.TypeOf(tab)
	t = t.Elem()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tab.TableName())
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
		if reflect.TypeOf(dbvalue).Kind() == reflect.Int {
			values = append(values, fmt.Sprintf("%d", dbvalue))
		} else {
			values = append(values, fmt.Sprintf("'%v'", dbvalue))
		}
	}

	sb.WriteString(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", t.TableName(), strings.Join(fields, ", "), strings.Join(values, ", ")))
	return sb.String()
}
type GoFakeitGenerator struct{}

type FakeDataGenerator interface {
	GenerateFakeUser() User
}

func (f *GoFakeitGenerator) GenerateFakeUser() User{
	u := User{
		ID: gofakeit.Number(1000, 9999),
		FirstName: gofakeit.FirstName(),
		LastName: gofakeit.LastName(),
		Email: gofakeit.Email(),
	}
	return u
}

func GenerateUserInserts(count int) []string {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	sqlFinal := []string{}
	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	sqlFinal = append(sqlFinal, sql)

	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		sqlFinal = append(sqlFinal, query)
	}
	return sqlFinal
}

func main() {
	// sqlGenerator := &SQLiteGenerator{}
	// fakeDataGenerator := &GoFakeitGenerator{}

	// user := User{}
	// sql := sqlGenerator.CreateTableSQL(&user)
	// fmt.Println(sql)

	// for i := 0; i < 34; i++ {
	// 	fakeUser := fakeDataGenerator.GenerateFakeUser()
	// 	query := sqlGenerator.CreateInsertSQL(&fakeUser)
	// 	fmt.Println(query)
	// }
	queries := GenerateUserInserts(34)

	// fmt.Println(queries)
	for _, query := range queries {
		fmt.Println(query)
	}
}
