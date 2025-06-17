package handlers

import (
	"encoding/json"
	"golang/weather/service"
	"log"
	"net/http"
)


type Handler struct{
	service service.Service
	responder Responder
}

type Bitcoin struct{
	USD float64
}

type Ethereum struct{
	USD float64
}

func NewHandler(service service.Service) *Handler{
	return &Handler{
		service: service,
		responder: NewResponder(),
	}
}

func (s *Handler) HandlerGetBitcoin(w http.ResponseWriter, r *http.Request){
	result, err := s.service.GetBitcoin(r.Context()) 
	if err != nil{
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	coin := Bitcoin{
		USD: result.Bitcoin.USD,
	}
	s.responder.OutputJSON(w, coin)
}

func (s *Handler) HandlerGetEthereum(w http.ResponseWriter, r *http.Request){
	result, err := s.service.GetEthereum(r.Context()) 
	if err != nil{
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	s.responder.OutputJSON(w, result)
}

type Respond struct{}

func NewResponder() Responder {
	return &Respond{}
}

type Responder interface{
	OutputJSON(w http.ResponseWriter, responseData interface{})
}

func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}){
	w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(responseData); err != nil {
        log.Printf("Ошибка кодирования в JSON: %v", err)
    }
}