package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)


// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:8080
// @BasePath  /
func main(){
	r := chi.NewRouter()

	r.Get("/api/login", handlerSearch)
	r.Get("/api/register", handlerGeocode)

	r.Post("/api/login", handlerLogin)
	r.Post("/api/register", handlerRegister)

	r.Group(func(r chi.Router){
		r.Use(jwtauth.Verifier(tokenJWT))
		r.Use(jwtauth.Authenticator)

		r.Post("/api/address/search", handlerSearch)
		r.Post("/api/address/geocode", handlerGeocode)
	})

	

	http.ListenAndServe(":8080", r)
}