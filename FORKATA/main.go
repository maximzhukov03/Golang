package main

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Username string
	Email    string
}



func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS users (ID SERIAL PRIMARY KEY, Username TEXT NOT NULL, Email INT NOT NULL);`

	_, err = db.Exec(query)
	if err != nil{
		return err
	}

	return nil
}

func InsertUser(user User) error {
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

func SelectUser(userID int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

	user := User{
		ID: userID,
	}
	query, args, err := PrepareQuery("select", "users", user)
	if err != nil{
		return user, err
	}
	row := db.QueryRow(query, args...)
	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil{
		return user, err
	}
	return user, nil
}

func UpdateUser(user User) error {
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

func DeleteUser(userID int) error {
	user := User{
		ID: userID,
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
		return sq.Insert(table).Columns("Username", "Email").Values(user.Username, user.Email).ToSql()

	case "update":
		return sq.Update(table).Set("Username", user.Username).Set("Email", user.Email).Where(squirrel.Eq{"ID": user.ID}).ToSql()

	case "delete":
		return sq.Delete(table).Where(squirrel.Eq{"ID": user.ID}).ToSql()

	case "select":
		return sq.Select("ID", "Username", "Email").From(table).Where(squirrel.Eq{"ID": user.ID}).ToSql()

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