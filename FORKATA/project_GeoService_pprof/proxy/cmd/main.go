package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "task25/proxy/docs"
	"task25/proxy/internal/auth"
	"task25/proxy/internal/controller"
	"task25/proxy/internal/repository"
	"task25/proxy/internal/reverse"
	"task25/proxy/internal/service"

	// address "task25/proxy/internal/service"
	"time"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:1313
// @BasePath  /
func main() {
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err)
    }
    userRepo := database.NewUserRepositoryPostgres(db)
    userSvc  := service.NewUserSevice(userRepo)

    apiKey, secret := os.Getenv("DA_API_KEY"), os.Getenv("DA_SECRET")
    geoSvc := service.NewGeoService(apiKey, secret)

    resp := controller.NewResponder()
    cont := controller.Controller{
        Service:   geoSvc,
        Responder: resp,
        Handler:   userSvc,
    }

    r := chi.NewRouter()
    rp := reverse.NewReverseProxy("hugo", "1313")

    r.Post("/auth/register", cont.HandlerRegister)
    r.Post("/auth/login",    cont.HandlerLogin)

    r.Get("/swagger/*", httpSwagger.Handler(
        httpSwagger.URL("http://localhost:1313/swagger/doc.json"),
    ))

    r.Use(rp.ReverseProxy)
    r.Group(func(r chi.Router) {
        r.Use(auth.RequireAuth)
        r.Route("/mycustompath/pprof", func(r chi.Router) {
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
        r.Post("/api/address/search",  cont.HandlerSearch)
        r.Post("/api/address/geocode", cont.HandlerGeocode)
    })

    fmt.Println("Starting server on :1313")
    log.Fatal(http.ListenAndServe(":1313", r))
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
