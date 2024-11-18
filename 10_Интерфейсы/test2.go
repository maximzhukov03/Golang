package main

import(
	"fmt"
)

func FindType(object interface{}){
	switch t := object.(type){
	case int: fmt.Println("int", t)
	case bool: fmt.Println("bool", t)
	case string: fmt.Println("string", t)
	case float32: fmt.Println("float", t)
	default: fmt.Println("Type is not found") 
	}
}

func test2(){
	FindType(123)
	FindType("qwerty")
	FindType(123.144234234)
	FindType(true)
	FindType(false)
}