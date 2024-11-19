package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct{
	ID int64
	Name string
	Email string
	Password string
	RegisterAt time.Time
}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=goLANG")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)

	}

	fmt.Println("CONECTED Наконец то блять")

	rows, err := db.Query("select * from users")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next(){
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisterAt)
		if err != nil{
			log.Fatal(err)
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(users)
}
