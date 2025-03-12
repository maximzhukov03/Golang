package main

import (
	"fmt"
	"strings"
)

func isRussianLetter(char rune) bool{
	if char >= 0x0400 && char <= 0x042F{
		return true
	}
	if char >= 0x0430 && char <= 0x044F{
		return true
	}
	return false
}

func countRussianLetters(s string) map[rune]int {
    counts := make(map[rune]int)
	txt := strings.ToLower(s)
    for _, char := range txt {
        if isRussianLetter(char) {
            counts[char] += 1
        }
    }
	
    return counts
}

func main() {
    result := countRussianLetters("Привет, мир!")
	for key, value := range result {
        fmt.Printf("%c: %d ", key, value) // в: 1 е: 1 т: 1 м: 1 п: 1 р: 2 и: 2 
    }
}