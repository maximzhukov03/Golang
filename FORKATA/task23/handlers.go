package main

import (
	"context"
	"encoding/json"
	"net/http"

	dadata "github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

var (
	tokenJWT = jwtauth.New("HS256", []byte("mysecretkey"), nil)
	users = map[string]string{}
	salt = "hj4h879ds8h3jklhf809sh"
)

type User struct{
	Name string `json:"name"`
	Password string `json:"password"`
}

type Address struct{
	Value string `json:"value"`
	City string `json:"city"`
}

type RequestAddressSearch struct {
  Query string `json:"query"`
}

type ResponseAddress struct {
  Addresses []*Address `json:"addresses"`
}

type RequestAddressGeocode struct {
  Lat string `json:"lat"`
  Lng string `json:"lng"`
}

// @Summary      Address Search
// @Description  get address
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        req  body RequestAddressSearch true  "Get address"
// @Success      200  {object}  ResponseAddress
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/address/search [post]
func handlerSearch(w http.ResponseWriter, r *http.Request){
	var req RequestAddressSearch
	var res ResponseAddress
	creds := client.Credentials{
		ApiKeyValue: ("a232f4a2ca9f02d604128a65496fd52f7f9f8857"),
		SecretKeyValue: ("f0369fd57cb509fec49697904ecc2d248d4eba9c"),
	}

	api := dadata.NewSuggestApi(client.WithCredentialProvider(&creds))
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	query := &suggest.RequestParams{Query: req.Query}
    result, err := api.Address(context.Background(), query)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, elem := range result{
		addr := &Address{
			Value: elem.Value,
			City: elem.Data.City,
		}
		res.Addresses = append(res.Addresses, addr)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

	
}

// @Summary      Address from Geocode
// @Description  get address from geocode
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        req  body RequestAddressGeocode true  "Get address from geocode"
// @Success      200  {object}  ResponseAddress
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/address/geocode [post]
func handlerGeocode(w http.ResponseWriter, r *http.Request){
	var req RequestAddressGeocode
	var res ResponseAddress
		creds := client.Credentials{
		ApiKeyValue: ("a232f4a2ca9f02d604128a65496fd52f7f9f8857"),
		SecretKeyValue: ("f0369fd57cb509fec49697904ecc2d248d4eba9c"),
	}

	api := dadata.NewSuggestApi(client.WithCredentialProvider(&creds))
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	query := &suggest.GeolocateParams{
		Lat: req.Lat,
		Lon: req.Lng,
	}
    result, err := api.GeoLocate(context.Background(), query)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, elem := range result{
		addr := &Address{
			Value: elem.Value,
			City: elem.Data.City,
		}
		res.Addresses = append(res.Addresses, addr)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// @Summary      Register user
// @Description  Регистрация нового пользователя. Хранение пароля осуществляется с использованием bcrypt.
// @Accept       json
// @Produce      json
// @Param        user  body  User  true  "User credentials"
// @Success      201  {string}  string  "User registered successfully"
// @Failure      400  {string}  string  "Invalid input or user already exists"
// @Failure      500  {string}  string  "Internal server error"
// @Router       /api/register [post]
func handlerRegister(w http.ResponseWriter, r *http.Request){
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, ok := users[user.Name]
	if ok{
		http.Error(w, "The user has already been added", http.StatusBadRequest)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password + salt), bcrypt.DefaultCost)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	users[user.Name]=string(hash)

}

// @Summary      Login user
// @Description  Аутентификация пользователя. Возвращает JWT токен при успешном входе.
// @Accept       json
// @Produce      json
// @Param        user  body  User  true  "User credentials"
// @Success      200  {object}  map[string]string  "JWT token"
// @Failure      400  {string}  string  "Invalid credentials or user not found"
// @Failure      500  {string}  string  "Internal server error"
// @Router       /api/login [post]
func handlerLogin(w http.ResponseWriter, r *http.Request){
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	passwd, ok := users[user.Name]
	if !ok{
		http.Error(w, "The user not found", http.StatusBadRequest)
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwd), []byte(user.Password + salt))
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, token, err := tokenJWT.Encode(map[string]interface{}{"sub": user.Name})
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

