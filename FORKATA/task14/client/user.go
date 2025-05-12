// Код клиента
package main

import (
	"fmt"
	"net"
)

func main() {
	// Устанавливаем соединение с сервером
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}
	defer conn.Close()

	// Отправка данных на сервер
	message := "Привет от клиента!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Ошибка при отправке данных:", err)
		return
	}

	// Чтение ответа от сервера
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
		return
	}

	// Вывод полученного ответа
	fmt.Printf("Получен ответ от сервера: %s\n", string(buffer[:n]))
}
