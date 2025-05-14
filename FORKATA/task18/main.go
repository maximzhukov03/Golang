package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func handleGroup1_1(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group 1 Привет, мир 1"))
}

func handleGroup1_2(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group 1 Привет, мир 2"))
}

func handleGroup1_3(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group 1 Привет, мир 3"))
}

func handleGroup2_1(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group 2 Привет, мир 1"))
}

func handleGroup2_2(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group 2 Привет, мир 2"))
}

func handleGroup2_3(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group 2 Привет, мир 3"))
}

func handleGroup3_1(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group 3 Привет, мир 1"))
}

func handleGroup3_2(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group 3 Привет, мир 2"))
}

func handleGroup3_3(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group 3 Привет, мир 3"))
}

func handleGroup2(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет, мир 2"))
}

func handleGroup3(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет, мир 3"))
}

func RouterConfigure() http.Handler{
	r := chi.NewRouter()
	r.Route("/group1", func(r chi.Router){
		r.Get("/1", handleGroup1_1)
		r.Get("/2", handleGroup1_2)
		r.Get("/3", handleGroup1_3)
	})
	r.Route("/group2", func(r chi.Router){
		r.Get("/", handleGroup2)
		r.Get("/1", handleGroup2_1)
		r.Get("/2", handleGroup2_2)
		r.Get("/3", handleGroup2_3)
	})
	r.Route("/group3", func(r chi.Router){
		r.Get("/", handleGroup3)
		r.Get("/1", handleGroup3_1)
		r.Get("/2", handleGroup3_2)
		r.Get("/3", handleGroup3_3)
	})
	return r

}



func main(){
	http.ListenAndServe(":8080", RouterConfigure())
}