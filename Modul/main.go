package main

import (
	"fmt"
	"golang-ninja/basic/shape"
)

func main(){
	Square := shape.NewSquare(5)
	printShapeArea(Square)
}

func printShapeArea(figure shape.Shape){ /*Создание функции для вывода Пощади фигуры*/
	fmt.Println(figure.Area())
}