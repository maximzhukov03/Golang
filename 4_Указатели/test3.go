/*
Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
*/
package main

import (
	"fmt"
)

func test3(){
	var x, y int = 5, 10
	summm(&x, &y)
	message := "STROKA KAK PRIMER" /*0x1234*/
	PrintMessage(&message) /*пеоедаем не значение а передаем ссылку область памяти*/
	fmt.Println(message) 
}

func summm(x *int, y *int){
	*x, *y = *y, *x
	fmt.Println(*x, *y)
}


func PrintMessage(message *string){ /*0x1235 но если мы добавим * то будем работать с ячейкой 0x1234 */
	*message += " DOPOLNENIE " /*переменная в другой ячейки памяти нежели оригинал*/
}