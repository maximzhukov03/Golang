/*Проверка значений есть ли в маппе или нет*/

package main

import "fmt"

func test2(){
	fmt.Printf("Запуск ТЕСТА №2\n")
	MyMap := map[string]int{"Num1":100, "Num2":200, "Num3":300}
	numInt, exist := MyMap["Num4"] /*exist принимает либо true либо false*/
	if exist{
		fmt.Println("Yes, I have Num1", numInt)
	}else{
		fmt.Println("NO", exist)
	}

}