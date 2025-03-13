package main

import (
	"encoding/json"
	"fmt"
)

func getJSON(data []User) (string, error) {
	dataJson, err := json.Marshal(data)
	if err != nil{
		return "", err
	}

	return string(dataJson), nil
}

type User struct {
    Name     string    `json:"name"`
    Age      int       `json:"age"`
    Comments []Comment `json:"comments"`
}

type Comment struct {
    Text string `json:"text"`
}

func main() {
	data := []User{
		{
			Name:  "Иван",
			Age:   30,
			Comments: []Comment{
				{Text: "Привет"},
				{Text: "Как дела"},
			},
		},

	}

	dataJson, err := getJSON(data)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(dataJson)
	
}