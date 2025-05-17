package main

import (
	"context"
	"fmt"
	"os"

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
		ApiKeyValue: os.Getenv("USER_API_KEY"),
		SecretKeyValue: os.Getenv("USER_API_KEY_SECRET"),
	}
	api := dadata.NewSuggestApi(client.WithCredentialProvider(&creds))
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := &suggest.RequestParams{Query: req.Query}
    res, err := api.Address(context.Background(), query)

	for _, elem := range res{
		addr := &Address{
			Value: elem.Value,
			City: elem.Data.City,
		}
		result.Addresses = append(result.Addresses, addr)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// func handlerGeocode(w http.ResponseWriter, r *http.Request){
// 	var req RequestAddressGeocode
// 	var result ResponseAddress
// 	apiHelp := dadata.NewSuggestApi()
// 	api := 
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	query := &suggest.R{Lat: req.Lat, Lon: req.Lng, }
//     res, err := api.Address(context.Background(), query)

// 	for _, elem := range res{
// 		addr := &Address{
// 			Value: elem.Value,
// 			City: elem.Data.City,
// 		}
// 		result.Addresses = append(result.Addresses, addr)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(result)
// }


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
	// r.Post("/api/address/geocode", handlerGeocode)
	http.ListenAndServe(":8080", r)
}