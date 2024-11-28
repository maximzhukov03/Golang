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
	email        sql.NullString // Изменено с *string на string
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
	
	err = InsertUser(db, User{name: "Wil", second_name: "Sirka", email: sql.NullString{String: "Sirka@mail.com", Valid: true}, date_of_birth: time.Date(2024, time.June, 1,0,0,0,0, time.UTC)})
	if err != nil {
		log.Fatal(err)
	}
	
	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, elem := range users {
		var email string
		if !elem.email.Valid{
			email = "НЕТ ПОЧТЫ"
		}else{
			email = elem.email.String
		}
		fmt.Printf("[ID]: %d| [Name]: %s %s, [email]: %s, [Date]: %s\n", elem.id, elem.name, elem.second_name, email, elem.date_of_birth.Format("2006-01-02"))
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
	var email string
	if !u.email.Valid{
		email = "НЕТ ПОЧТЫ"
	}else{
		email = u.email.String
	}
	_, err := db.Exec("INSERT INTO employee (name, second_name, email) VALUES ($1, $2, $3)", u.name, u.second_name, email)
	return err
}