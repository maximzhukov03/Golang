package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func generateActivationKey() string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())
	passwordList := []string{}
	var b strings.Builder
	for j := 0; j < 4; j++{
		for i := 0; i < 4; i++ {
			b.WriteRune(rune(chars[rand.Intn(len(chars))]))
			fmt.Println(b.String())
		}
		passwordList = append(passwordList, b.String())
		b = strings.Builder{}
	}
	return strings.Join(passwordList, "-")
}

func main() {
	activationKey := generateActivationKey()
	fmt.Println(activationKey) // UQNI-NYSI-ZVYB-ZEFQ
}