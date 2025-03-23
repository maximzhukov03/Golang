package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Структура для парсинга ответа API Binance
type PriceResponse struct {
	Coin string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

func getPrice(Coin string) (float64, error) {
	var priceResponse PriceResponse
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", Coin)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		return 0, err
	}

	return priceResponse.Price, nil
}

func main() {
	btcUsdt := []float64{}
	ethUsdt := []float64{}
	ltcUsdt := []float64{}

	for i := 0; i < 5; i++{
		time.Sleep(1 * time.Second)	
		btcPrice, err := getPrice("BTCUSDT")
		if err != nil {
			log.Fatalf("Error fetching BTC price: %v", err)
		}
		btcUsdt = append(btcUsdt, btcPrice)
	
		ltcPrice, err := getPrice("LTCUSDT")
		if err != nil {
			log.Fatalf("Error fetching LTC price: %v", err)
		}
		ltcUsdt = append(ltcUsdt, ltcPrice)

	
		ethPrice, err := getPrice("ETHUSDT")
		if err != nil {
			log.Fatalf("Error fetching ETH price: %v", err)
		}
		ethUsdt = append(ethUsdt, ethPrice)
	}



	fmt.Printf("BTC-USD: %.2f\n", btcUsdt)
	fmt.Printf("LTC-USD: %.2f\n", ltcUsdt)
	fmt.Printf("ETH-USD: %.2f\n", ethUsdt)
}