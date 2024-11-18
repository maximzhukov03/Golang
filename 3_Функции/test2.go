/*
Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.

Входные данные
Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).

Выходные данные
Необходимо вернуть значение функции от x, y и z.
*/

package main

import (
	"fmt"
)

func test2(){
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	Golosovanie(x, y, z)
}
func Golosovanie(x, y, z int){
	if x + y + z > 1{
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}