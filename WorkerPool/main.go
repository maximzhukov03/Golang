package main

import (
	"fmt"
	"time"
)

func main(){
	times := time.Now()
	const jobsCount, workerCount = 15, 3
	jobs := make(chan int, jobsCount)
	result := make(chan int, jobsCount)

	for i := 0; i < jobsCount; i++{
		jobs <- i + 1
	}
	close(jobs)
	for i := 0; i < workerCount; i++{
		go worker(i+1, jobs, result)		
	}
	

	for i := 0; i < jobsCount; i++{
		fmt.Printf("result #%d : value = %d\n", i+1, <- result)
	}

	fmt.Println("ВРЕМЕНИ ПРОШЛО: ", time.Since(times).String())
}


func worker(id int, jobs <-chan int, result chan<- int){
	for j := range jobs{
		time.Sleep(time.Second)
		fmt.Printf("worker #%d finished\n", id)
		result <- j * j
	}
}


// func main(){
// 	const jobsCount, workerCount = 15, 3
// 	times := time.Now()

// 	jobs := make(chan int, jobsCount)
// 	result := make(chan int, jobsCount)

// 	for i := 0; i < workerCount; i++{
// 		go worker(i+1, jobs, result)
// 	}
// 	go worker(1, jobs, result)

// 	for i := 0; i < jobsCount; i++{
// 		jobs <- i + 1
// 	}
// 	close(jobs)

// 	for i := 0; i < jobsCount; i++{
// 		fmt.Printf("result #%d : value = %d\n", i+1, <- result)
// 	}

// 	fmt.Println("Времени прошло: ", time.Since(times).String())
// }

// //                только для чтения    запись и чтение
// func worker(id int, jobs <-chan int, result chan<- int) {
// 	for j := range jobs {
// 		time.Sleep(time.Second)
// 		fmt.Printf("worker #%d finished\n", id)
// 		result <- j*j
// 	}
// }