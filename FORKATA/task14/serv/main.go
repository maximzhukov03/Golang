// Код сервера
package main

import (
	"fmt"
	"net"
)

func main() {
	// Устанавливаем прослушивание на порту 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен. Ожидание подключений...")

	// Принимаем входящие подключения
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии подключения:", err)
			continue
		}

		// Обработка подключения в отдельной горутине
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Чтение данных от клиента
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
		return
	}

	// Вывод полученных данных
	fmt.Printf("Получены данные от клиента: %s\n", string(buffer[:n]))

	// Отправка ответа клиенту
	response := "Привет от сервера!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Ошибка при отправке данных:", err)
		return
	}

	fmt.Println("Ответ успешно отправлен клиенту.")
}