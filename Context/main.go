package main

import (
	"context"
	"fmt"
	"time"
)

//  context.TODO - когда не уверенны какой контекст использовать
// context.Value - стоит использовать как можно реже и передавать только необязательные параметры
// ctx всегда передается первым аргументом в функцию
// нельзя передавать пустой контекст (nil)
// та функция которая, инициализировала контекст, может ее отменить
func main(){
	ctx := context.Background() // Только на самом высоком уровне (Инициализируем функцией маин)
	ctx, _ = context.WithTimeout(ctx, time.Second * 3)	// Задаем сколько секунд позволительно висеть

	parse(ctx)
}

func parse(ctx context.Context){
	for{
		select{
		case <- time.After(time.Second * 2) :
			fmt.Println("Parsing complited")
			return 
		case <- ctx.Done():
			fmt.Println("Deadline exceded")
			return
		}

	}
}