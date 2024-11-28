package main

import (
	"database/sql"
	//"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type User struct {
	id         int64
	name       string
	second_name      string
	email   *string
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

	rows, err := db.Query("select * from employee")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.id, &u.name, &u.second_name, &u.email, &u.date_of_birth)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// var us User
	// err = db.QueryRow("select id, name from employee where id = $1", 2).Scan(&us.id, &us.name)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		fmt.Println("no rows")
	// 		return
	// 	}
	// 	log.Fatal(err)
	// }

	// fmt.Println(us)
	for _, elem := range users{
		email := "nil"
		if elem.email != nil{
			email = *elem.email
		}
		fmt.Printf("[ID]: %d, [Name]: %s %s, [email]: %s, [Date]: %s\n", elem.id, elem.name, elem.second_name, email, elem.date_of_birth.Format("2006-01-02"))
	}
}
