package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error{
			fmt.Println(req.Response.Status)
			fmt.Println("ReDireCTED")
			return nil
		},
	}
	resp, err := client.Get("http://google.com")
	if err != nil {
		log.Fatal(err) //Логировать и закрывать приложение
	}

	defer resp.Body.Close() // У ответа есть тело Body

	fmt.Println("Response status: ", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
