// // написать функцию, котораю мержит 4 канала в один
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func merge(ch1, ch2, ch3, ch4 chan int) chan int{
//     wg := sync.WaitGroup{}
//     chMerge := make(chan int)
//     wg.Add(4)
//     go func(){
//         defer wg.Done()
//         for elem := range ch1{
//             chMerge <- elem
//         }
//     }()
//     go func(){
//         defer wg.Done()
//         for elem := range ch2{
//             chMerge <- elem
//         }
//     }()
//     go func(){
//         defer wg.Done()
//         for elem := range ch3{
//             chMerge <- elem
//         }
//     }()
//     go func(){
//         defer wg.Done()
//         for elem := range ch4{
//             chMerge <- elem
//         }
//     }()
//     go func() {
//         wg.Wait()
//         close(chMerge)
//     }()
//     return chMerge
// }

// func main(){
//     ch1 := make(chan int)
//     ch2 := make(chan int)
//     ch3 := make(chan int)
//     ch4 := make(chan int)

//     go func(){
//         for i := 1; i < 6; i++{
//             ch1 <- i
//         }  
//         close(ch1)
//     }()
//     go func(){
//         for i := 6; i < 11; i++{
//             ch2 <- i
//         } 
//         close(ch2)
//     }()
//     go func(){
//         for i := 11; i < 16; i++{
//             ch3 <- i
//         }  
//         close(ch3)
//     }()
//     go func(){
//         for i := 16; i < 21; i++{
//             ch4 <- i
//         } 
//         close(ch4)
//     }()
    
    
//     ch := merge(ch1, ch2, ch3, ch4)

//     for elem := range ch{
//         fmt.Println(elem)
//     }

// }

// // написать функцию которая пишет в канал N количество элементов
// package main

// import (
// 	"fmt"
// )

// func writer(ch chan interface{}, list []interface{}){
//     for _, elem := range list{
//         ch <- elem
//     }
//     close(ch)
// }

// func main(){
//     ch := make(chan interface{})

//     list := []interface{}{1, 2, "Привет", true}
    
//     go writer(ch, list)

//     for elem := range ch{
//         fmt.Println(elem)
//     }
// }


// // написать функцию которая пишет в канал N количество элементов

// package main

// import "fmt"

// func writer(ch chan int, N int){
//     for i := 0; i < N; i++{
//         ch <- i
//     }
//     close(ch)
// }

// func main(){
//     ch := make(chan int)
    
//     go writer(ch, 5)

//     for elem := range ch{
//         fmt.Println(elem)
//     }
// }


// // принимает список чисел, вычисляет их квадраты параллельно в нескольких горутинах и отправляет результаты в выходной канал.

// package main

// import "fmt"

// func square(ch, ch2 chan int){
//     for elem := range ch{
//         ch2 <- elem * elem
//     }
//     close(ch2)
// }


// func main(){
//     listInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//     chan1 := make(chan int)
//     chan2 := make(chan int)
//     go func(){
//         for _, elem := range listInt{
//             chan1 <- elem
//         }
//         defer close(chan1)
//     }()

//     go square(chan1, chan2)

//     for elem := range chan2{
//         fmt.Println(elem)
//     }
// }


// Работа со строкой из двух каналов по подсчету символов в строке
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	chan1 := make(chan string)
// 	chan2 := make(chan int)
//     lines := []string{"Я","Ты","Она",}
// 	go func() {
// 		for str := range chan1 {
// 			count := len([]rune(str))
// 			chan2 <- count
// 		}
// 		close(chan2)
// 	}()

	

// 	go func() {
// 		for _, line := range lines {
// 			chan1 <- line
// 		}
// 		close(chan1)
// 	}()

// 	for count := range chan2 {
// 		fmt.Println(count)
// 	}
// }

// #########################################################################################
// #########################################################################################
// #########################################################################################
// #########################################################################################
// #########################################################################################
// #########################################################################################
// #########################################################################################
// #########################################################################################


// package main

// import ()


// type A struct {
// 	a int
// }
// type B struct {
// 	a A
// }

// func main(){
// 	a := B.a.a
// }


// // Задание 5: Напишите функцию, которая подсчитывает количество уникальных слов в строке

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"testing"
// )

// func wordCount(s string) map[string]int {
//     strSlise := strings.Fields(s)
// 	mapString := make(map[string]int)
// 	for _, elem := range strSlise{
// 		mapString[elem]--
// 	}
// 	return mapString
// }

// func TestWordCount(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected map[string]int
// 	}{
// 		{"Hello world Hello", map[string]int{"Hello": 2, "world": 1}},
// 		{"Go is great Go", map[string]int{"Go": 2, "is": 1, "great": 1}},
// 		{"", map[string]int{}}, // Пустая строка
// 		{"a a a a", map[string]int{"a": 4}}, // Повторение одного слова
// 		{"one two three four", map[string]int{"one": 1, "two": 1, "three": 1, "four": 1}}, // Все уникальные слова
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.input, func(t *testing.T) {
// 			result := wordCount(tt.input)
// 			for key, expectedValue := range tt.expected {
// 				if result[key] != expectedValue {
// 					t.Errorf("wordCount(%s) = %v; expected %v", tt.input, result, tt.expected)
// 				}
// 			}
// 		})
// 	}
// }


// func main(){
// 	str := "Hello world Hello"
// 	mapStirng := wordCount(str)
// 	for elem, count := range mapStirng{
// 		fmt.Printf("Слов: %s | Колличество: %d\n", elem, count)
// 	}

// }



// // Задание 4: Напишите функцию, которая удаляет все элементы
// // с заданным значением из среза.
// // Функция должна возвращать новый срез, не изменяя исходный.

// package main

// import "fmt"

// func removeElement(arr []int, value int) []int {
// 	arrNew := []int{}
// 	for _, elem := range arr{
// 		if elem != value{
// 			arrNew = append(arrNew, elem)
// 		}	
// 	}
// 	return arrNew
// }

// func main(){
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	result := removeElement(arr, 3)
// 	for _, elem := range result{
// 		fmt.Println(elem)
// 	}
// }


// // Задача 3: Напишите функцию, которая принимает строку
// // и возвращает карту, где ключи — это символы из строки,
// // а значения — количество их вхождений.

// package main

// import "fmt"

// func countCharacterFrequency(str string) map[rune]int{
// 	runes := []rune(str)
// 	mapStr := make(map[rune]int)
// 	for _, elem := range runes{
// 		mapStr[elem]++
// 	}
// 	return mapStr
// }

// func main(){
// 	str := "hello world"
// 	result := countCharacterFrequency(str)
// 	for elem, count := range result{
// 		fmt.Printf("|Символ: %s | Колличество: %d|\n", string(elem), count)
// 	}

// }











// // Задача 2: Напишите функцию, которая принимает срез целых чисел
// // и удаляет все четные числа из этого среза.
// // Функция должна возвращать новый срез без четных чисел.

// package main

// import "fmt"

// func removeEvenNumbers(arr []int) []int{
// 	newArr := []int{}
// 	for _, elem := range arr{
// 		if (elem % 2) != 0{
// 			newArr = append(newArr, elem)
// 		}
// 	}
// 	return newArr
// }

// func main(){
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	result := removeEvenNumbers(arr)
// 	fmt.Println(result)
// }








// // Задача 1: Написать функцию выводящую дублирующиеся значения в
// // слайсе из заданного массива. Вывод должен быть []int
// // и максимальное значение.

// package main

// import "fmt"

// func findDoubleAndMaxVals(arr *[]int) ([]int, int){
// 	arrValues := make(map[int]int)
// 	arrValuesSlice := []int{}
// 	maximalka := (*arr)[0]
// 	for _, elem := range *arr{
// 		arrValues[elem]++
// 		if elem > maximalka{
// 			maximalka = elem
// 		}
// 	}
// 	for key, value := range arrValues{
// 		if value > 1{
// 			arrValuesSlice = append(arrValuesSlice, key)
// 		}
// 	}
// 	return arrValuesSlice, maximalka

// }


// func main(){
// 	arr := []int{1, 2, 2, 3, 3, 4, 5, 5, 7, 6, 1, 10}

// 	doubleValue, maxVal := findDoubleAndMaxVals(&arr)
// 	for _, elem := range doubleValue{
// 		fmt.Printf("Повторяющееся число: %d\n", elem)
// 	}
// 	fmt.Printf("Максимальное число: %d", maxVal)

// }