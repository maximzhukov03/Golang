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
	ID    int    `json:"id"`
	FIRST_NAME  string `json:"first_name"`
	SECOND_NAME string `json:"second_name"`
	EMAIL string `json:"email"`
	PASSWORD string `json:"password"`
}

var users []User

func main(){

	connect := "host=127.0.0.1 port=5432 user=postgres dbname=users_log sslmode=disable password=goLANG"
	db, err := sql.Open("postgres", connect)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil{
		log.Fatal(err)
	}
	users, err = GetUsers(db)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("CONECTED")

	http.HandleFunc("/user", handleUser)
	http.ListenAndServe("localhost:8080", nil)

	// users, err := GetUsers(db)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(users)

}

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func GetUsers(db *sql.DB) ([]User, error){
	rows, err := db.Query("SELECT * FROM user_data")
	if err != nil{
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next(){
		u := User{}
		err := rows.Scan(&u.ID, &u.FIRST_NAME, &u.SECOND_NAME, &u.EMAIL, &u.PASSWORD)
		if err != nil{
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil{
		return nil, err
	}

	return users, nil
}

func GetUser(db *sql.DB, id int) ([]User, error){
	rows, err := db.Query("SELECT * FROM user_data where id = $1", id)
	if err != nil{
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next(){
		u := User{}
		err := rows.Scan(&u.ID, &u.FIRST_NAME, &u.SECOND_NAME, &u.EMAIL, &u.PASSWORD)
		if err != nil{
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil{
		return nil, err
	}

	return users, nil
}