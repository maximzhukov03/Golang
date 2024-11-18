package main

import(
	"fmt"
)

func FindTypeOK(object interface{}){
	v, ok := object.(string)
	if ok{
		fmt.Println("Это строка", v)
	} else {
		fmt.Println("это не строка")
	}
}

func test3(){
	a1 := 12
	a2 := "STRING"

	FindTypeOK(a1)
	FindTypeOK(a2)
}