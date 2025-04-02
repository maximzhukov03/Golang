package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

func (u *User) TableName() string {
	return "users"
}

type Tabler interface {
	TableName() string
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

type SQLiteGenerator struct{}

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
		values = append(values, fmt.Sprintf("'%v'", dbvalue))
	}

	sb.WriteString(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", t.TableName(), strings.Join(fields, ", "), strings.Join(values, ", ")))
	return sb.String()
}

type Migrator struct{
	db *sql.DB
	sqlGenerator SQLGenerator
}


func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	return &Migrator{
		db:           db,
		sqlGenerator: sqlGenerator,
	}
}

func (m Migrator) Migrate(models ...Tabler) error{
	for _, model := range models{
		tableSQL := m.sqlGenerator.CreateTableSQL(model)

		_, err := m.db.Exec(tableSQL)
		if err != nil{
			return fmt.Errorf("failed to create table for model %v: %v", model.TableName(), err)
		}
	}
	return nil
}

// Основная функция
func main() {
	// Подключение к SQLite БД
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	YourSQLGeneratorInstance := &SQLiteGenerator{}
	// Создание мигратора с использованием вашего SQLGenerator
	migrator := NewMigrator(db, YourSQLGeneratorInstance)

	// Миграция таблицы User
	if err := migrator.Migrate(&User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}