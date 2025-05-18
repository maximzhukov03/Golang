package main

import (
	"context"
	"fmt"

	"encoding/json"
	"net/http"

	dadata "github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-chi/chi"
)

type Address struct{
	Value string
	City string
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


func handlerSearch(w http.ResponseWriter, r *http.Request){
	var req RequestAddressSearch
	var result ResponseAddress
	creds := client.Credentials{
		ApiKeyValue: ("a232f4a2ca9f02d604128a65496fd52f7f9f8857"),
		SecretKeyValue: ("f0369fd57cb509fec49697904ecc2d248d4eba9c"),
	}
	api := dadata.NewSuggestApi(client.WithCredentialProvider(&creds))
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(req)
	query := &suggest.RequestParams{Query: req.Query}
    res, err := api.Address(context.Background(), query)
	if err != nil{
		http.Error(w, "DaData" + err.Error(), http.StatusInternalServerError)
	}
	for _, elem := range res{
		addr := &Address{
			Value: elem.Value,
			City: elem.Data.City,
		}
		result.Addresses = append(result.Addresses, addr)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func handlerGeocode(w http.ResponseWriter, r *http.Request){

	var req RequestAddressGeocode
	var result ResponseAddress
	creds := client.Credentials{
		ApiKeyValue: ("a232f4a2ca9f02d604128a65496fd52f7f9f8857"),
		SecretKeyValue: ("f0369fd57cb509fec49697904ecc2d248d4eba9c"),
	}
	api := dadata.NewSuggestApi(client.WithCredentialProvider(&creds))
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := &suggest.GeolocateParams{Lat: req.Lat, Lon: req.Lng, }
    res, err := api.GeoLocate(context.Background(), query)
	if err != nil{
		http.Error(w, "DaData" + err.Error(), http.StatusInternalServerError)
		return
	}
	for _, elem := range res{
		addr := &Address{
			Value: elem.Value,
			City: elem.Data.City,
		}
		result.Addresses = append(result.Addresses, addr)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}


func DaDataExample()  {
	api := dadata.NewSuggestApi()

	params := suggest.RequestParams{
		Query: "ул Свободы",
	}

	suggestions, err := api.Address(context.Background(), &params)
	if err != nil {
		return
	}

	for _, s := range suggestions {
		fmt.Printf("%s", s.Value)
	}
}

func main(){
	r := chi.NewRouter()
	r.Post("/api/address/search", handlerSearch)
	r.Post("/api/address/geocode", handlerGeocode)
	http.ListenAndServe(":8080", r)
}