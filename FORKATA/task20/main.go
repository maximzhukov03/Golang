// Пример кода для создания сервера на go-chi

package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func RouterConfigure(r *chi.Mux) http.Handler{
	r.Use(LoggerMiddleware)
	r.Get("/hello", handleHello)
	return r
}

func Logger(){
	logger, _ = zap.NewProduction()
}

var logger *zap.Logger

func main() {
	r := chi.NewRouter()

	// Применение middleware для логирования с помощью zap logger
	Logger()
	defer logger.Sync()

	// Здесь можно добавить ваши маршруты с различными методами
	RouterConfigure(r)
	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(
			"CONNECT",
  			zap.String("PATH:", r.Method),
  			zap.String("URL: ", r.URL.Path),
  			zap.String("ID: ", r.RemoteAddr),
		)
		next.ServeHTTP(w, r)
	})
}

func handleHello(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}