package main

import (
	"fmt"
	"time"
)

func main(){
	message1 := make(chan string)
	message2 := make(chan string)

	go func(){
		for{
			message1 <- "Канал 1 работает 200мс"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func(){
		for{
			message2 <- "Канал 2 работает 1с"
			time.Sleep(time.Second * 1)
		}
	}()

	for{
		select{
		case msg := <-message1: fmt.Println(msg) 
		case msg := <-message2: fmt.Println(msg)
		default:
		}		
	}
}