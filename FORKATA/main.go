package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}


func CreateUserTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT NOT NULL, age INT NOT NULL);`

	_, err := db.Exec(query)
	if err != nil{
		return err
	}

	return nil
}

func InsertUser(db *sql.DB, user User) error {
	query := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err := db.Exec(query, user)
	if err != nil{
		return err
	}
	return nil
}

func SelectUser(db *sql.DB, id int) (User, error) {
	var user User
	query := `SELECT id, name, age FROM users WHERE id = ?`
	row := db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil{
		return user, err
	}
	return user, nil
}

func UpdateUser(db *sql.DB, user User) error {
	query := `UPDATE users SET name = ?, age = ? WHERE id = ?`
	_, err := db.Exec(query, &user.Name, &user.Age, &user.ID)
	if err != nil{
		return err
	}
	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil{
		return err
	}
	return nil
}

func main() {
    db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()
}