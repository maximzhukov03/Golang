package main

import(
	"fmt"
)

func test1(){
	fmt.Printf("Запуск ТЕСТА №1\n")
	users := map[int]int{1: 101, 2:201, 3:301, 4:401, 5:501} /*ключ значение*/
	for i := 1; i <= len(users); i++{
		fmt.Println(users[i])
	}
}