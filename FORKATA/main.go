package main

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}


func CreateUserTable(db *sql.DB) error {
	db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT NOT NULL, age INT NOT NULL);`

	_, err = db.Exec(query)
	if err != nil{
		return err
	}

	return nil
}

func InsertUser(db *sql.DB, user User) error {
	db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	query, args, err := PrepareQuery("insert", "users", user)
	if err != nil{
		return err
	}

	_, err = db.Exec(query, args...)
	if err != nil{
		return err
	}
	return nil
}

func SelectUser(db *sql.DB, id int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	var user User
	query, args, err := PrepareQuery("select", "users", user)
	if err != nil{
		return user, err
	}
	row := db.QueryRow(query, args...)
	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil{
		return user, err
	}
	return user, nil
}

func UpdateUser(db *sql.DB, user User) error {
	db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	query, args, err := PrepareQuery("select", "users", user)
	if err != nil{
		return err
	}
	_, err = db.Exec(query, args...)
	if err != nil{
		return err
	}
	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	user := User{
		ID: id,
	}
	db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	query, args, err := PrepareQuery("select", "users", user)
	if err != nil{
		return err
	}
	_, err = db.Exec(query, args...)
	if err != nil{
		return err
	}
	return nil
}

func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)

	switch operation {
	case "insert":
		return sq.Insert(table).Columns("name", "age").Values(user.Name, user.Age).ToSql()

	case "update":
		return sq.Update(table).Set("name", user.Name).Set("age", user.Age).Where(squirrel.Eq{"id": user.ID}).ToSql()

	case "delete":
		return sq.Delete(table).Where(squirrel.Eq{"id": user.ID}).ToSql()

	case "select":
		return sq.Select("id", "name", "age").From(table).Where(squirrel.Eq{"id": user.ID}).ToSql()

	default:
		return "", nil, fmt.Errorf("неизвестная операция: %s", operation)
	}
}

func main() {
    db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()
}
