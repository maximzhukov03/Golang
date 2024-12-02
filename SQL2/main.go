package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main(){
	connect := "host=127.0.0.1 port=5432 user=postgres dbname=Users sslmode=disable password=goLANG"
	db, err := sql.Open("postgres", connect)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil{
		log.Fatal(err)
	}

	fmt.Println("CONECTED")
}
