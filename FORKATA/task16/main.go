package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil{
		return
		fmt.Println("Ошибка: ", err)
	}
	partsOfLines := strings.Split(strings.TrimSpace(line), " ")
	if partsOfLines[0] == "GET" && partsOfLines[1] == "/"{
		response := "HTTP/1.1 200 OK\n" +
			"Content-Type: text/html\n" +
			"\n" +
			"<!DOCTYPE html>\n" +
			"<html>\n" +
			"<head>\n" +
			"<title>Webserver</title>\n" +
			"</head>\n" +
			"<body>\n" +
			"hello world\n" +
			"</body>\n" +
			"</html>\n"
		conn.Write([]byte(response))
	} else {
		response := "HTTP/1.1 404 Not Found\n" +
			"Content-Type: text/plain\n" +
			"\n" +
			"404 page not found\n"
		conn.Write([]byte(response))
	}


}

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}