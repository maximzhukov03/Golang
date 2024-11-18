/*
Неограниченый ввод в функцию и анонимная функция
*/

package main

import "fmt"

func main(){
	a := FindMin(1, 2, 43, 45, 323, 545, 1324,  -3, -143)
	fmt.Println(a)

	func (){
		fmt.Println("Hello") /*анонимная функция*/
	}()

	inc := increment()

	fmt.Println(inc()) /*1*/
	fmt.Println(inc()) /*2*/
	fmt.Println(inc()) /*3*/
	fmt.Println(inc()) /*4*/
}
func FindMin(numbers ...int) int{ /*Неограниченый ввод*/
	if len(numbers) == 0{
		return 0
	}

	min := numbers[0]

	for _, elem := range numbers{
		if elem < min{
			min = elem
		}
	}
	return min
}

func increment() func () int{
	count := 0

	return func() int{
		count++
		return count
	}
}