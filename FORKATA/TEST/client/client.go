package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type List struct{
	list []int
}


func main(){
	resp, err := http.Get("http://localhost:8080/ping")
	if err != nil{
		log.Println("ОШИБКА В ПОЛУЧЕНИИ ОТВЕТА КЛИЕНТОМ")
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Println("ОШИБКА В ПОЛУЧЕНИИ ОТВЕТА КЛИЕНТОМ")
	}
	num, err := strconv.Atoi(string(b))
	if err != nil{
		log.Println("ОШИБКА В ПЕРЕВОДЕ ЧИСЛА В СТРОКУ", err)
	}
	fmt.Println(calc(num))
	

}

func calc(num int) int{
	return num * num
}