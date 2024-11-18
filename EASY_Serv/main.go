package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


type Person struct{ //Структура для профиля
	Name string
	Age string
}

var people []Person // Массив чтобы можно было хранить все профили (Здесь была бы БД)

func main(){
	http.HandleFunc("/person", peopleHandler) // функция которая связывает определенный URL-путь с функцией-обработчиком.:
	http.HandleFunc("/info", healthCheckHandler) // Она принимает два аргумента Путь и Обработчик (handler):
	//функция, которая будет вызвана при получении запроса по указанному пути. Сигнатура func(http.ResponseWriter, *http.Request)

	log.Println("server start listen")
	err := http.ListenAndServe("localhost:8080", nil) //  запускает HTTP-сервер, который слушает входящие запросы на указанном сетевом адресе и порту.
	if err != nil{
		log.Fatal(err)
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request){ // В r хранится тело запроса (заголовки куки и т.д)
	switch r.Method{
		case http.MethodGet: getPeople(w, r)
		case http.MethodPost: postPeople(w, r)
		default: http.Error(w,"Invalid HTTP method", http.StatusMethodNotAllowed)
	} 
}

func getPeople(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(people)
	fmt.Fprintf(w, "get people %v", people)
}

func postPeople(w http.ResponseWriter, r *http.Request){
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	people = append(people, person)
	fmt.Fprintf(w, "post new person %v", person)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "http was correctly")
}