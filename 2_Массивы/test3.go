/*
Описание
Создайте программу, которая реверсирует (переворачивает) массив целых чисел.

Условия
Создайте массив целых чисел с произвольными значениями.
Напишите функцию reverseArray, которая принимает массив и реверсирует его.
Выведите исходный и реверсированный массивы на экран.
*/
package main

import (
	"fmt"
)


func test3(){
	mas := [5]int{}
	for i := 0; i < len(mas); i++{ fmt.Scan(&mas[i]) }
	masReverse := [len(mas)]int{}
    for i := range mas { masReverse[i] = mas[len(mas)-1-i] }
	fmt.Println("Перевернутый массив: ", masReverse)
}