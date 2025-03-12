package main

import (
	"fmt"
	"regexp"
)

func main() {
	email := "test@example.com"
	valid := isValidEmail(email)
	if valid {
		fmt.Printf("%s является валидным email-адресом\n", email)
	} else {
		fmt.Printf("%s не является валидным email-адресом\n", email)
	}
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`[@]`)
	re2 := regexp.MustCompile("\\.com")
	matches := re.MatchString(email)
	matches2 := re2.MatchString(email)
	if matches && matches2{
		return true
	} else {
		return false
	}
}