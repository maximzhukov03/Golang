package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	myPanic(ch)

	fmt.Println(<-ch)
}

func myPanic(ch chan string) {
	panic("my panic message")
}