package main

import (
	//"fmt"
	"time"
)

func trySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false

	}

}

func main() {
	ch := make(chan int, 1)
	trySend(ch, 48)
	time.Sleep(1*time.Second)
}