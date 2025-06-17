package main

import (
	"fmt"
	"golang/weather/controller"
	"golang/weather/service"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func main(){
	r := chi.NewRouter()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Ошибка соединения с Redis:", err)
		return
	}
	fmt.Println("Соединение с Redis успешно:", pong)
	service := service.NewService(client)
	handler := handlers.NewHandler(service)
	r.Get("/bitcoin", handler.HandlerGetBitcoin)
	r.Get("/ethereum", handler.HandlerGetEthereum)
	
	log.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}


// type ServiceGet interface{
// 	GetEthereum()
// 	GetBitcoin()
// }
