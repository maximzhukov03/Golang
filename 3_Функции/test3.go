/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:
a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/

package main

import "fmt"

func test3(){
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println("Сложение чисел: ", a, b, c, d)
	count, sum := sumInt(a, b, c, d)
	fmt.Printf("Всего чисел: %d\n", count)
	fmt.Printf("Сумма чисел: %d", sum)
}

func sumInt(args ...int) (int, int){
	sum := 0
	for _, elem := range args{ sum += elem }
	return len(args), sum
}