/*Создание структуры*/
package main

import (
	"fmt"
)

type userNew struct{
	Name string
	Age int
	Sex string
	Height int
}


func test1(){
	fmt.Printf("ЗАПУСК ТЕСТА №1\n")
	NewUser := userNew{"Максимка", 22, "Мужичек", 198}
	user := struct{
		Name string
		Age int
		Sex string
		Height int
	}{"Максим", 21, "Мужик", 195}
	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", NewUser)
}