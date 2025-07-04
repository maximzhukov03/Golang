package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/pprof"
	"task25/proxy/internal/service"
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
	Service service.GeoProvider
	Responder Responder
}

// @Summary      Address Search
// @Description  get address
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

// @Summary Pprof Index
// @Description Все доступные pprof-профили
// @Tags diagnostics
// @Produce text/html
// @Success 200 {string} string "HTML со списком профилей"
// @Security BearerAuth
// @Router /mycustompath/pprof/index [get]
func PprofIndex(w http.ResponseWriter, r *http.Request) {
    pprof.Index(w, r)
}

// @Summary Pprof Cmdline
// @Description cmdline
// @Tags diagnostics
// @Produce text/html
// @Success 200 {string} string "HTML со списком профилей"
// @Security BearerAuth
// @Router /mycustompath/pprof/cmdline [get]
func PprofCmdline(w http.ResponseWriter, r *http.Request){
	pprof.Cmdline(w, r)
}

// @Summary CPU-профиль
// @Description Снимает CPU-профиль за заданное количество секунд (параметр seconds)
// @Tags diagnostics
// @Param seconds query int false "Продолжительность профилирования в секундах" default(30)
// @Produce application/octet-stream
// @Success 200 {file} file "pprof CPU binary"
// @Security BearerAuth
// @Router /mycustompath/pprof/profile [get]
func PprofProfile(w http.ResponseWriter, r *http.Request) {
    pprof.Profile(w, r)
}

// @Summary Symbol Lookup
// @Description Выполняет преобразование адресов в имена символов (symbol lookup)
// @Tags diagnostics
// @Param symbol query string true "Адрес или символ для поиска"
// @Produce application/octet-stream
// @Success 200 {file} file "pprof symbol binary"
// @Security BearerAuth
// @Router /mycustompath/pprof/symbol [get]
func PprofSymbol(w http.ResponseWriter, r *http.Request) {
    pprof.Symbol(w, r)
}

// @Summary Runtime Trace
// @Description Снимает трассировку работы рантайма за заданное количество секунд (параметр seconds)
// @Tags diagnostics
// @Param seconds query int false "Длительность трассировки в секундах" default(1)
// @Produce application/octet-stream
// @Success 200 {file} file "pprof trace binary"
// @Security BearerAuth
// @Router /mycustompath/pprof/trace [get]
func PprofTrace(w http.ResponseWriter, r *http.Request) {
    pprof.Trace(w, r)
}

// @Summary Allocation Samples
// @Description Профиль выборочных аллокаций памяти (allocs)
// @Tags diagnostics
// @Produce application/octet-stream
// @Success 200 {file} file "pprof allocs binary"
// @Security BearerAuth
// @Router /mycustompath/pprof/allocs [get]
func PprofAllocs(w http.ResponseWriter, r *http.Request) {
    pprof.Handler("allocs").ServeHTTP(w, r)
}

// @Summary Block Profile
// @Description Профиль ожиданий блокировок (block profile)
// @Tags diagnostics
// @Produce application/octet-stream
// @Success 200 {file} file "pprof block binary"
// @Security BearerAuth
// @Router /mycustompath/pprof/block [get]
func PprofBlock(w http.ResponseWriter, r *http.Request) {
    pprof.Handler("block").ServeHTTP(w, r)
}

// @Summary Goroutine Snapshot
// @Description Снимок текущих горутин (goroutine profile)
// @Tags diagnostics
// @Produce application/octet-stream
// @Success 200 {file} file "pprof goroutine binary"
// @Security BearerAuth
// @Router /mycustompath/pprof/goroutine [get]
func PprofGoroutine(w http.ResponseWriter, r *http.Request) {
    pprof.Handler("goroutine").ServeHTTP(w, r)
}

// @Summary Heap Profile
// @Description Снимок кучи (heap profile)
// @Tags diagnostics
// @Produce application/octet-stream
// @Success 200 {file} file "pprof heap binary"
// @Security BearerAuth
// @Router /mycustompath/pprof/heap [get]
func PprofHeap(w http.ResponseWriter, r *http.Request) {
    pprof.Handler("heap").ServeHTTP(w, r)
}

// @Summary Mutex Profile
// @Description Профиль ожидания мьютексов (mutex profile)
// @Tags diagnostics
// @Produce application/octet-stream
// @Success 200 {file} file "pprof mutex binary"
// @Security BearerAuth
// @Router /mycustompath/pprof/mutex [get]
func PprofMutex(w http.ResponseWriter, r *http.Request) {
    pprof.Handler("mutex").ServeHTTP(w, r)
}

// @Summary Thread Creation Profile
// @Description Профиль создания системных потоков (threadcreate profile)
// @Tags diagnostics
// @Produce application/octet-stream
// @Success 200 {file} file "pprof threadcreate binary"
// @Security BearerAuth
// @Router /mycustompath/pprof/threadcreate [get]
func PprofThreadcreate(w http.ResponseWriter, r *http.Request) {
    pprof.Handler("threadcreate").ServeHTTP(w, r)
}

