package main

import (
	"fmt"
	"golang/httpclient/coincap"
	"log"
	"time"
)

func main(){
	coincapClient, err := coincap.NewClient(time.Second * 2)
	if err != nil{
		log.Fatal(err)
	}

	assets, err := coincapClient.GetAssets()
	if err != nil{
		log.Fatal(err)
	}

	for _, asseassets := range assets{
		fmt.Println(asseassets.Info())
	}

	botcoin, err := coincapClient.GetAsset("bitcoin")
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(botcoin.Info())
}