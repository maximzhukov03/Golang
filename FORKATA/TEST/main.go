package main

import (
	"encoding/json"
	"net/http"
)

type List struct{
	list []int
}

func main(){
	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "pong")
	// })

	http.HandleFunc("/ping", JSONhandler)

	http.ListenAndServe(":8080", nil)
}

func JSONhandler(w http.ResponseWriter, r *http.Request){
	list :=  List{
		list: []int{1, 2, 3, 4},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(list)
	if err != nil{
		return
	}

}