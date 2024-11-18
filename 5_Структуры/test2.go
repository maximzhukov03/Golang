/*Конструкторы*/

package main

import (
	"fmt"
)

type Usering struct{
	Name string
	Age int
	Sex string
	Height int
}


func ConstructUser(name, sex string, age, height int) Usering{
	return Usering{
		Name: name,
		Age: age,
		Sex: sex, 
		Height: height,
	}
}

func test2(){
	fmt.Printf("ЗАПУСК ТЕСТА №2\n")
	a := ConstructUser("Макисм", "Мужик", 21, 195)
	fmt.Println(a)
}