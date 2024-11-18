/*Создание интерфейса*/
/*Если у структуры есть метод (в данном случае Area()) то проходит*/

package main

import(
	"fmt"
)

type Shape interface{ /*Создание интерфейса*/
	Area() float32
}

type Square struct{ /*Создание структуры Квадрат*/
	sideLen float32
}
 
func (s Square) Area() float32{ /*Создание функции с методом Area() для подсчета площади квадрата*/
	return s.sideLen * s.sideLen
}

func printShapeArea(figure Shape){ /*Создание функции для вывода Пощади фигуры*/
	fmt.Println(figure.Area())
}

func test1(){
	Square := Square{5}

	printShapeArea(Square)
}