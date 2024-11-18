/*пробежка по Мапе и удаление объектов*/

package main

import(
	"fmt"
)

func test3(){
	fmt.Printf("Запуск ТЕСТА №3\n")
	MyNewMap := map[string]int{"GEORG": 4444, "JORJE": 5555, "VAHTANG": 6666}
	for key, elem := range MyNewMap{
		fmt.Println(key, elem)
	}
	delete(MyNewMap, "GEORG")
	fmt.Println("Маппа после удаления")
	for key, elem := range MyNewMap{
		fmt.Println(key, elem)
	}
}