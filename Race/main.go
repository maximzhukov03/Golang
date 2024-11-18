package main

import (
	"fmt"
	"sync"
)

type Bank struct{
	balance int
	mu *sync.Mutex
}

func (b *Bank) DEPOSIT(money int) *Bank{
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += money
	return b
}

func (b *Bank) TAKE(money int) *Bank{
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance -= money
	return b
}

func OPERATION(b *Bank, wg *sync.WaitGroup){
	b.DEPOSIT(1000)
	b.TAKE(400)
	wg.Done()
}

func main(){

	wg := &sync.WaitGroup{}

	b := Bank{
		balance: 10000,
		mu: new(sync.Mutex),
	}

	for i := 0; i < 1000; i++{
		wg.Add(1)
		go OPERATION(&b, wg)
	}
	wg.Wait()
	fmt.Println(b.balance)
}