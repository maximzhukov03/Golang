package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"task25/proxy/internal/controller"
	"task25/proxy/internal/reverse"
	address "task25/proxy/internal/service"
	_ "task25/proxy/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"time"
	"github.com/go-chi/chi"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:1313
// @BasePath  /
func main() {
	r := chi.NewRouter()
	reverseProxy := reverse.NewReverseProxy("hugo", "1313")
	apiKey := "a232f4a2ca9f02d604128a65496fd52f7f9f8857"
	secretKey := "f0369fd57cb509fec49697904ecc2d248d4eba9c"
	resp := controller.NewResponder()
	service := address.NewGeoService(apiKey, secretKey)

	cont := controller.Controller{
		Responder: resp,
		Service: service,
	}

	r.Use(reverseProxy.ReverseProxy)
	r.Get("/swagger/*", httpSwagger.Handler())
	r.Get("/api/address/search", cont.HandlerSearch)
	r.Get("/api/address/geocode ", cont.HandlerGeocode)

	r.Route("/mycustompath", func(r chi.Router) {
		r.Get("/pprof/index", controller.PprofIndex)
		r.Get("/pprof/cmdline", controller.PprofCmdline)
		r.Get("/pprof/profile", controller.PprofProfile)
		r.Get("/pprof/symbol", controller.PprofSymbol)
		r.Get("/pprof/trace", controller.PprofTrace)
		r.Get("/pprof/allocs", controller.PprofAllocs)
		r.Get("/pprof/block", controller.PprofBlock)
		r.Get("/pprof/goroutine", controller.PprofGoroutine)
		r.Get("/pprof/heap", controller.PprofHeap)
		r.Get("/pprof/mutex", controller.PprofMutex)
		r.Get("/pprof/threadcreate", controller.PprofThreadcreate)
	})
	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", r)
}

const content = ``

func WorkerTest() {
	t := time.NewTicker(1 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
