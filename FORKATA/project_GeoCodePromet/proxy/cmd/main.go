package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "task25/proxy/docs"
	"task25/proxy/internal/auth"
	"task25/proxy/internal/config"
	"task25/proxy/internal/controller"
	"task25/proxy/internal/repository"
	"task25/proxy/internal/reverse"
	"task25/proxy/internal/service"
	"time"
	_ "github.com/lib/pq"
	"github.com/go-redis/redis"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host      localhost:8080
// @BasePath  /
func main() {
	c := config.NewConfig()
	conf := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		c.DB_HOST, c.DB_PORT, c.DB_USER, c.DB_PASSWORD, c.DB_NAME)
	
	db, err := sql.Open("postgres", conf)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Пинг БД: %v", err)
	}
    userRepo := database.NewUserRepositoryPostgres(db)
    userSvc  := service.NewUserSevice(userRepo)
	apiKey := "a232f4a2ca9f02d604128a65496fd52f7f9f8857"
	secretKey := "f0369fd57cb509fec49697904ecc2d248d4eba9c"
    geoSvc := service.NewGeoService(apiKey, secretKey)

    resp := controller.NewResponder()
    cont := controller.Controller{
        Service:   geoSvc,
        Responder: resp,
        Handler:   userSvc,
    }

    r := chi.NewRouter()
    rp := reverse.NewReverseProxy("localhost", "1313")
	r.Use(rp.ReverseProxy)
    r.Post("/auth/register", cont.HandlerRegister)
    r.Post("/auth/login",    cont.HandlerLogin)

    r.Get("/swagger/*", httpSwagger.Handler(
        httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
    ))


    r.Group(func(r chi.Router) {
        r.Use(auth.RequireAuth)
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
        r.Post("/api/address/search",  cont.HandlerSearch)
        r.Post("/api/address/geocode", cont.HandlerGeocode)
    })

    fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
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
