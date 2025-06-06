package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	listen, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Println("ОШИБКА В СОЕДИНЕНИИ")
	}
	defer listen.Close()

	for{
		conn, err := listen.Accept()
		if err != nil{
			continue
		}
		go send(conn)
	}
}

func send(conn net.Conn){
	defer conn.Close()
	for i := 1; i <= 4; i++{
		time.Sleep(time.Second * 1)
		_, err := fmt.Fprint(conn, i)
		if err != nil{
			log.Println("ОШИБКА в отправке")
			return
		}
	}
}