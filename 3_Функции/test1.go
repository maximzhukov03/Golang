/*
Напишите функцию f(), которая будет принимать строку text и выводить ее (печатать на экране).
*/
package main

import (
	"fmt"
)

func test1(){
	var str string
	fmt.Println("Введите слово: ")
	fmt.Scan(&str)
	f(str)
}

func f(s string){
	fmt.Println("Вы ввели слово: ", s)
}