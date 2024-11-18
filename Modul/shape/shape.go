package shape

import(
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

func NewSquare(length float32) Square{
	return Square{
		sideLen: length,
	}
}



