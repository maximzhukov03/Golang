package main

import (
	"fmt"
	"sync"
)

func waitGroupExample(goroutines ...func() string) string {
	var wg sync.WaitGroup
	res := make(chan string, len(goroutines))
	for _, goRout := range goroutines{
		wg.Add(1)
		go func(goRout func() string){
			defer wg.Done()
			res <- goRout()
		}(goRout)
	}
	wg.Wait()
	close(res)

	var stringOutput string
	for i := range res{
		stringOutput += fmt.Sprintf("%s\n", i)
	}
	return stringOutput
}

func main() {
	count := 1000
	goroutines := make([]func() string, count)

	for i := 0; i < count; i++ {
		j := i
		goroutines[i] = func() string {
			return fmt.Sprintf("goroutine %d", j)
		}
	}

	fmt.Println(waitGroupExample(goroutines...))
}