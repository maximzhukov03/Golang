package main

import (
	"fmt"
	"runtime"
	"time"
)

func factorialRecursive(n int) int {
	if n <= 1{
		return 1
	}
	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	if n <= 1{
		return 1
	}
	count := 1
	for i := 1; i <= n; i++{
		count  *= i
	}
	return count
}

// выдает true, если реализация быстрее и false, если медленнее
func compareWhichFactorialIsFaster() map[string]bool {
	start := time.Now()
	factorialIterative(100000)
	res1 := time.Since(start)

	start = time.Now()
	factorialRecursive(100000)
	res2 := time.Since(start)
	
	res := map[string]bool{
		"factorialIterative": res1 < res2,
		"factorialRecursive": res2 < res1,
	}
	return res	
}

func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("Go OS/Arch:", runtime.GOOS, "/", runtime.GOARCH)

	fmt.Println("Which factorial is faster?")
	fmt.Println(compareWhichFactorialIsFaster())
}