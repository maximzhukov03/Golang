/*Дженерики 
Они позволяют разработчикам писать функции и структуры данных, 
которые могут работать с различными типами без необходимости создавать 
отдельные версии для каждого типа. Это особенно полезно для уменьшения 
дублирования кода и повышения его читаемости.*/
package main

import(
	"fmt"
)

type Number interface{
	int | float64
}

func main(){
	a := []int{1, 2, 3, 4}
	b := []float64{1.1, 2.2, 3.3, 4.5}
	c := []string{"1", "2", "3"}
	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(sumNumberInterface(a))
	fmt.Println(sumNumberInterface(b))
	fmt.Println(SearchMass(c, "2"))
}

func sum[v int | float64](input []v) v{
	var result v
	for _, elem := range input{
		result += elem
	}
	return result
}

func sumNumberInterface[v Number](input []v) v{
	var result v
	for _, elem := range input{
		result += elem
	}
	return result
}

func SearchMass[C comparable](elements []C, searchEl C) bool{ /*Вместо comporable */
	for _, el := range elements{/*								использовать можно Any*/
		if el == searchEl{
			return true
		} else {

		}
	}
	return false
}