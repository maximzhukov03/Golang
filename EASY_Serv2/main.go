package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct{
	ID int `json:"id"`
	NAME string `json:"name,omitempty"`
}

var (
	users = []User{{1, "Vasya"}, {2, "Petya"}}
)

func main(){
	http.HandleFunc("/user", loggerMiddleware(handleUser))
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Info"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Index page")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "index.html") // чтобы отправить страницу используется ServeFile
	})
	http.ListenAndServe("localhost:8080", nil)
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		idFromCtx := r.Context().Value("id")
		userID, ok := idFromCtx.(string)
		if !ok{
			log.Printf("[%s] - %s - error: UserID is invalid\n", r.Method, r.URL)
			w.WriteHeader(http.StatusBadRequest)
		}

		log.Printf("[%s] - %s - by userID - [%s]\n", r.Method, r.URL, userID)
		next(w, r)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet: getUser(w, r)
	case http.MethodPost: addUser(w, r)
	default: w.WriteHeader(405)
	}
}

func getUser(w http.ResponseWriter, r *http.Request){
	resp, err := json.Marshal(users)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func addUser(w http.ResponseWriter, r *http.Request){
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user User
	if err = json.Unmarshal(reqBytes, &user); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users = append(users, user)
}

// import (
// 	"context"
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"net/http"
// )

// type User struct{
// 	ID int `json:"id"`
// 	NAME string `json:"name,omitempty"`
// }

// var (
// 	users = []User{{1,"Vasya"}, {2,"Petya"}}
// )

// func main(){
// 	http.HandleFunc("/user", authMiddleware(loggerMiddleware(handleUser)))

// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil{
// 		log.Fatal(err)
// 	}	
// }

// func authMiddleware(next http.HandlerFunc) http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request){
// 		userID := r.Header.Get("x-id")
// 		if userID == "" {
// 			log.Printf("[%s] %s - error: UserID is not provided\n", r.Method, r.RequestURI)
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		ctx := r.Context()

// 		ctx = context.WithValue(ctx, "id", userID)

// 		r = r.WithContext(ctx)

// 		next(w,r)
// 	}

	

// }

// func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request){
// 		idFromCtx := r.Context().Value("id")
// 		userID, ok := idFromCtx.(string)
// 		if !ok{
// 			log.Printf("[%s] %s - error: UserID is invalid\n", r.Method, r.URL)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		log.Printf("[%s] %s by userID %s\n", r.Method, r.URL, userID)
// 		next(w, r)
// 	}

// }

// func handleUser(w http.ResponseWriter, r *http.Request){
// 	switch r.Method{
// 	case http.MethodGet: getUser(w, r)
// 	case http.MethodPost: addUser(w, r)
// 	default: w.WriteHeader(http.StatusNotImplemented)
// 	}	
// }

// func getUser(w http.ResponseWriter, r *http.Request){
// 	resp, err := json.Marshal(users)
// 	if err != nil{
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write(resp)
// }

// func addUser(w http.ResponseWriter, r *http.Request){
// 	reqBytes, err := io.ReadAll(r.Body)
// 	if err != nil{
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	var user User
// 	if err = json.Unmarshal(reqBytes, &user); err != nil{
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	users = append(users, user)
// }