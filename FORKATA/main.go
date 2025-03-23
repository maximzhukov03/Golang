package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		time.Sleep(1 * time.Second)
		fmt.Print("\033[H\033[2J") // очистка терминала
		timeNow := time.Now().Format("15:04:05")
		timeDate := time.Now().Format("2006-01-02")
		// dateNow := time.Date()
		fmt.Printf("%s\n%s", timeNow, timeDate)
	}
}