/*
Описание
Напишите программу, которая удаляет элемент из среза по заданному индексу.

Условия
Создайте срез строк с именами: "Alice", "Bob", "Charlie", "David".
Напишите функцию remove, которая принимает срез и индекс элемента для удаления.
После удаления элемента выведите обновленный срез на экран.
*/

package main

import "fmt"

func test2(){
	sr1 := []string{}
	sr1 = append(sr1, "Alise", "Bob", "Charlie", "David")
	var srIndex int
	fmt.Println("\nВведите индекс слова для удаления: ")
	fmt.Scan(&srIndex)
	nameDel := sr1[srIndex]
	sr1 = append(sr1[0:srIndex], sr1[srIndex+1:] ...)
	fmt.Println("Удалено слово: ", nameDel)
	fmt.Println("Преобразованный срез", sr1)
}