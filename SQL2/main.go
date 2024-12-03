package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
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
var user User
func main(){
	http.HandleFunc("/user", handleUser)
	http.ListenAndServe("localhost:8080", nil)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		connect := "host=127.0.0.1 port=5432 user=postgres dbname=users_log sslmode=disable password=goLANG"
		db, err := sql.Open("postgres", connect)
		if err != nil{
			log.Fatal(err)
		}
		defer db.Close()
	
		if err := db.Ping(); err != nil{
			log.Fatal(err)
		}
		log.Println("CONECTED")
		users, err = GetUsers(db)
		if err != nil{
			log.Fatal(err)
		}
		log.Println("CONECTED")
		getUser(w, r)
	case http.MethodPost:
		connect := "host=127.0.0.1 port=5432 user=postgres dbname=users_log sslmode=disable password=goLANG"
		db, err := sql.Open("postgres", connect)
		if err != nil{
			log.Fatal(err)
		}
		defer db.Close()
	
		if err := db.Ping(); err != nil{
			log.Fatal(err)
		}
		log.Println("CONECTED")
		TakeUser(w, r)
		fmt.Println(user)
		err = InsertUser(db, user)
		if err != nil{
			log.Fatal(err)
		}
	case http.MethodDelete:
		connect := "host=127.0.0.1 port=5432 user=postgres dbname=users_log sslmode=disable password=goLANG"
		db, err := sql.Open("postgres", connect)
		if err != nil{
			log.Fatal(err)
		}
		defer db.Close()
	
		if err := db.Ping(); err != nil{
			log.Fatal(err)
		}
		log.Println("CONECTED")
		TakeUser(w, r)
		err = DeleteUser(db, user)
		if err != nil{
			log.Fatal("NOT DELETE")
		}


	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) { // Вывод на сервер Юзера
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func TakeUser(w http.ResponseWriter, r *http.Request) User{ // Размещение юзера через сервер
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
	}
	if err = json.Unmarshal(reqBytes, &user); err != nil{
		w.WriteHeader(http.StatusBadRequest)
	}
	return user
}

func GetUsers(db *sql.DB) ([]User, error){ // Получение ВСЕХ Юзеров в БД
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
 
func GetUser(db *sql.DB, id int) ([]User, error){ //  Получение Юзера в БД
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

func InsertUser(db *sql.DB, u User) error { // Добавление Юзера в БД
	_, err := db.Exec("INSERT INTO user_data (first_name, second_name, email, password) VALUES ($1, $2, $3, $4)", u.FIRST_NAME, u.SECOND_NAME, u.EMAIL, u.PASSWORD)
	if err != nil{
		log.Fatal("NOT INSERT USER")
		return err
	}
	log.Println("ADDED USER")
	return err
}

func DeleteUser(db *sql.DB, u User) error {
	_, err := db.Exec("DELETE FROM user_data where first_name = $1", u.FIRST_NAME)
	log.Println("DETETE USER")
	return err
}