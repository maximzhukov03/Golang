package main

import (
	"fmt"
	"os"
)

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
    file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, perm)
	if err != nil{
		return err
	}
	file.WriteString(string(data))
	return nil
}


func main(){
	err := WriteFile("/path/to/file.txt", []byte("Hello, World!"), os.FileMode(0644))
	if err != nil {
	    fmt.Println(err)
		return
	}
}