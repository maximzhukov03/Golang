package main

import (
	//"fmt"
	"fmt"
	"time"
)

func timeout(timeout time.Duration) func() bool{
	ch := make(chan bool, 1)
	
	go func(){
		time.Sleep(timeout)
		ch <- true
	}()

	return func() bool{
		select{
		case <- ch:
			return true
		case <- time.After(timeout):
			return false
		}
	}
}
// Пример использования функции timeout
func main() {
    timeoutFunc := timeout(3 * time.Second)
    since := time.NewTimer(3050 * time.Millisecond)
    for {
        select {
        case <-since.C:
            fmt.Println("Функция не выполнена вовремя")
            return
        default:
            if timeoutFunc() {
                fmt.Println("Функция выполнена вовремя")
                return
            }
        }
    }
}
// Output: Функция выполнена вовремя