package main

import (
	"fmt"
)

func countUniqueUTF8Chars(s string) int {
    rune := make(map[rune]bool)
	for _, elem := range s{
		rune[elem] = true
	}
	return len(rune)
}

func main(){
	a := "Hello, 世界!"
	fmt.Println(countUniqueUTF8Chars(a))
}