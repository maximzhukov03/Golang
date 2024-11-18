package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	NAME  string `json:"name,omitempty"`
	EMAIL string `json:"email,omitempty"`
}

var users = []User{{1, "Petya", "maxmaxiiim2@mal.ruu"}, {2, "Gosha", "wqeaxiiim2@mal.r"}}

func main() {
	http.HandleFunc("/user", handleUser)
	http.ListenAndServe("localhost:8080", nil)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPost:
		postUser(w, r)
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

func postUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users = append(users, user)
}
