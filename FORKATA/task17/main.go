package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Hello world")
}

func main(){
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}