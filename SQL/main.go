package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	id           int64
	name         string
	second_name  string
	email        string // Изменено с *string на string
	date_of_birth time.Time
}

func main() {
	connect := "host=127.0.0.1 port=5432 user=postgres dbname=Users sslmode=disable password=goLANG"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("CONECTED Наконец то блять")

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, elem := range users {
		email := "[НЕ ИМЕЕТ]"
		if elem.email != "" { // Проверка на пустую строку вместо nil
			email = elem.email
		}
		fmt.Printf("[ID]: %d| [Name]: %s %s, [email]: %s, [Date]: %s\n", elem.id, elem.name, elem.second_name, email, elem.date_of_birth.Format("2006-01-02"))
	}

	err = InsertUser(db, User{name: "William", second_name: "Sir", email: "Sir@mail.com"})
	if err != nil {
		log.Fatal(err)
	}
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.id, &u.name, &u.second_name, &u.email, &u.date_of_birth)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func InsertUser(db *sql.DB, u User) error {
	_, err := db.Exec("INSERT INTO employee (name, second_name, email) VALUES ($1, $2, $3)", u.name, u.second_name, u.email)
	return err
}