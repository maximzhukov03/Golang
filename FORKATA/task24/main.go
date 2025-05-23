// Пример кода для реализации API-сервера с graceful shutdown
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Создание HTTP-сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux, // Здесь должен быть ваш обработчик запросов
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	finalSignal := make(chan struct{})
	sysCallChan := make(chan os.Signal, 1)
	signal.Notify(sysCallChan, os.Interrupt, syscall.SIGTERM)
	
	http.HandleFunc("/test", handler)

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	go func(){
		// Ожидание сигнала остановки
		signal := <- sysCallChan
		log.Printf("Сигнал: %s пришел", signal)
		// Создание контекста с таймаутом для graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Остановка сервера с использованием graceful shutdown
		err := server.Shutdown(ctx)
		if err != nil{
			log.Println(err)
		}
		close(finalSignal)
	}()

	<-finalSignal
	log.Println("Server stopped gracefully")
}


func handler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	log.Println("ОБРАБОТЧИК РАБОТАЕТ")
	w.Write([]byte("Тестовая обработка запроса"))
}