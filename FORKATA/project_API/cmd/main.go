package main

import(
	"github.com/go-chi/chi"
)

func main(){
	r := chi.NewRouter()
	r.Get("/api/users/{id}", handlerGetId)
}