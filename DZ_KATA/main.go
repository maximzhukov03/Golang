package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Text struct {
	Content string
}

func (t *Text) textModifier() {
	t.Content = strings.Join(strings.Fields(t.Content), " ")
	// fmt.Println(t.Content)
	strHelp := []rune("")
	runString := []rune(t.Content)
	for i := 0; i < len(runString); i++ {
		if runString[i] == '-' {
			if i > 0 && i < len(runString)-1 {
				strHelp[len(strHelp)-1], runString[i+1] = runString[i+1], strHelp[len(strHelp)-1]
			}
			continue
		} else {
			strHelp = append(strHelp, runString[i])
		}
	}
	t.Content = string(strHelp)
	// fmt.Println(t.Content)

	t.Content = strings.ReplaceAll(t.Content, "+", "!")
	// fmt.Println(t.Content)

	srezHelp := []int{}
	strHelp = []rune("")
	runString = []rune(t.Content)
	for i := 0; i < len(runString); i++ {
		if unicode.IsDigit(runString[i]) {
			intNum, err := strconv.Atoi(string(runString[i]))
			if err != nil {
				fmt.Println(err)
			}
			srezHelp = append(srezHelp, intNum)
		} else {
			strHelp = append(strHelp, runString[i])
		}
	}
	t.Content = strings.Join(strings.Fields(string(strHelp)), " ")
	sum := 0

	for _, num := range srezHelp {
		sum += num
	}

	if sum > 0 {
		t.Content = fmt.Sprintf("%s %d", t.Content, sum)
		fmt.Println(t.Content)
	} else {
		fmt.Println(t.Content)
	}
}

func main() {
	text := &Text{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите строку: ")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}
}
