package main

import (
	"context"
	// "database/sql"
	"encoding/json"
	"fmt"
	"movie/golang/pkg/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	// _ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	prdlogger, _ := zap.NewProduction()
	defer prdlogger.Sync()
	logger := prdlogger.Sugar()

	logger.Infow("server started")

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("fataled to parse config %s", err)
	}

	fmt.Printf("cfg = %+v\n", cfg)
	// connect := "host=127.0.0.1 port=5432 user=postgres dbname=users_log sslmode=disable password=goLANG"
	// db, err := sql.Open("postgres", connect)
	// if err != nil{
	// 	logger.Fatal(err)
	// }
	// defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	srv := &http.Server{
		Handler: r,
		Addr:    cfg.HttpAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatalw("server failed", "err", err, "http-addr", cfg.HttpAddr)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	<-signals

	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalw("failed to shutdown server", "err", err, "http-addr", cfg.HttpAddr)
	}
	logger.Infow("buy!")
	// shutdown := make(chan struct{})
	// go func(){
	// 	signals := make(chan os.Signal, 1)
	// 	signal.Notify(signals, syscall.SIGTERM,syscall.SIGINT)
	// 	<-signals
	// 	close(shutdown)
	// }()

	// <-shutdown
	// Ожидание сигнала завершения
}
