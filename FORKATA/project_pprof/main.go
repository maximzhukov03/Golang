package main

import (
	"log"
	"runtime/pprof"
	"os"
)

func work(){
	f := 13
	for i:=0; i < 100000000; i++{
		_ = i * f
	}
	log.Printf("ОБРАБОТЧИК ЗАКОНЧИЛ РАБОТУ")  
}


func main(){
	file, err := os.Create("file.pprof")
	if err != nil{
		log.Printf("Ошибка в создание файла: %v", err)
	}
	defer file.Close()

	err = pprof.StartCPUProfile(file)
	if err != nil{
		log.Printf("Ошибка pprof: %v", err)
	}

	defer pprof.StopCPUProfile()

	work()
}