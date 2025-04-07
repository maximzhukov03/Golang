package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	// "strings"
)

type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo(opts ...func(exmo *Exmo)) *Exmo{
	exmo := &Exmo{
		client: http.DefaultClient,
		url: "https://api.exmo.com/v1.1",
	}
	for _, opt := range opts{
		opt(exmo)
	}

	return exmo
} 

func WithClient(client *http.Client) func(exmo *Exmo){
	return func(ex *Exmo){
		ex.client = client
	}
} 

func WithURL(url string) func(exmo *Exmo){
	return func(ex *Exmo){
		ex.url = url
	}
}

type Currencies []string

func GetConv(constanta string, url url.Values, exmo *Exmo) ([]byte, error){
	client := exmo.client
	urlReq := exmo.url + constanta
	if url != nil {
		urlReq += "?" + url.Encode()
	}
	req, err := http.NewRequest("GET", urlReq, nil)
  
	if err != nil {
	  return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
	  return nil, err
	}
	defer res.Body.Close()
  
	body, err := io.ReadAll(res.Body)
	if err != nil {
	  return nil, err
	}
	return body, nil
}

func (e *Exmo) GetCurrencies(){
	var curr Currencies
	
	body, err := GetConv("/currency", nil, e)
  if err != nil{
    fmt.Errorf("Ошибка")
  }


	err = json.Unmarshal(body, &curr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(curr)
}

func main(){
  exchange := NewExmo()

  exchange.GetCurrencies()
}
