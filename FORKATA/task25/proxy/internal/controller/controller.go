package controller

import (
	"errors"
	"log"
	"net/http"
	"task25/proxy/internal/usecase"
	"encoding/json"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Respond struct {
}

type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

func NewResponder() Responder {
	return &Respond{}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println("responder json encode error")
	}
}

func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	log.Println("http response bad request status code")
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		log.Println("response writer error on write")
	}
}

func (r *Respond) ErrorInternal(w http.ResponseWriter, err error) {
	log.Println("http response internal server error:", err)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: err.Error(),
	}); err != nil {
		log.Println("response writer error on write")
	}
}

type Controller struct{
	Service address.GeoProvider
	Responder Responder
}

// @Summary      Address Search
// @Description  get address
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        query  query string true  "Get address"
// @Success      200  {object}  Response
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/address/search [post]
func (c *Controller) HandlerSearch(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("query")
	if query == ""{
		log.Println("Проблема с обработкой query")
		c.Responder.ErrorBadRequest(w, errors.New("query parameter is required"))
		return
	}

	result, err := c.Service.AddressSearch(query)
	if err != nil{
		log.Println("Ошибка в AddressSearch")
		c.Responder.ErrorInternal(w, err)
		return
	}

	c.Responder.OutputJSON(w, Response{
		Success: true,
		Data: result,
	})
}

// @Summary      Address from Geocode
// @Description  get address from geocode
// @Accept       json
// @Produce      json
// @Param        query  query string true  "Get address from geocode"
// @Success      200  {object}  Response
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/address/geocode [post]
func (c *Controller) HandlerGeocode(w http.ResponseWriter, r *http.Request){
	lat := r.URL.Query().Get("lat")
	lng := r.URL.Query().Get("lon")
	if lat == "" || lng == ""{
		log.Println("Проблема с обработкой query")
		c.Responder.ErrorBadRequest(w, errors.New("query parameter is required"))
		return
	}

	result, err := c.Service.GeoCode(lat, lng)
	if err != nil{
		log.Println("Ошибка в AddressSearch")
		c.Responder.ErrorInternal(w, err)
		return
	}

	c.Responder.OutputJSON(w, Response{
		Success: true,
		Data: result,
	})
}

