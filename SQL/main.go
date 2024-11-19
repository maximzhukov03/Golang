package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/lib/pq"
)

func main(){

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=goLANG")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)


		fmt.Println("CONECTED")
	}
}