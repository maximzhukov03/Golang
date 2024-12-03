package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct{
	id int64
	first_name string
	second_name string
	email string
	password string
}

var users = []User{{1, "Mika", "Pedlo", "masdm@maisl.ru", "kjdsakflksjad3j3i2j4o"}}


func main(){

	http.HandleFunc("/user", handleUser)
	http.ListenAndServe("localhost:8080", nil)

	connect := "host=127.0.0.1 port=5432 user=postgres dbname=users_log sslmode=disable password=goLANG"
	db, err := sql.Open("postgres", connect)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil{
		log.Fatal(err)
	}

	fmt.Println("CONECTED")

	// users, err := GetUsers(db)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(users)

}

func handleUser(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet:
		GetUser(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request){
	resp, err := json.Marshal(users)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

// func GetUsers(db *sql.DB) ([]User, error){
// 	rows, err := db.Query("SELECT * FROM user_data")
// 	if err != nil{
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	users := make([]User, 0)
// 	for rows.Next(){
// 		u := User{}
// 		err := rows.Scan(&u.id, &u.first_name, &u.second_name, &u.email, &u.password)
// 		if err != nil{
// 			return nil, err
// 		}
// 		users = append(users, u)
// 	}

// 	err = rows.Err()
// 	if err != nil{
// 		return nil, err
// 	}

// 	return users, nil
// }

// func GetUser(db *sql.DB, id int) ([]User, error){
// 	rows, err := db.Query("SELECT * FROM user_data where id = $1", id)
// 	if err != nil{
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	users := make([]User, 0)
// 	for rows.Next(){
// 		u := User{}
// 		err := rows.Scan(&u.id, &u.first_name, &u.second_name, &u.email, &u.password)
// 		if err != nil{
// 			return nil, err
// 		}
// 		users = append(users, u)
// 	}

// 	err = rows.Err()
// 	if err != nil{
// 		return nil, err
// 	}

// 	return users, nil
// }