package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("ОШИБКА В ПОДКЛЮЧЕНИИ")
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	buffer := make([]byte, 1)

	for {
		_, err := reader.Read(buffer)
		if err != nil {
			return
		}


		char := buffer[0]

		num, err  := strconv.Atoi(string(char))
		if err != nil{
			log.Println("ОШИБКА В ПЕРЕВОДЕ СИМВОЛА")
		}

		fmt.Printf("Число: %d\n", num*2)
	}
}