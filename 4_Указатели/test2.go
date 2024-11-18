/*
Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.
*/

package main

import "fmt"

func test2(){
	var d, c int
	fmt.Scan(&d, &c)
	fmt.Println(summ(&d, &c))
}

func summ(x *int, y *int) int{
	i := *x + *y
	return i
}