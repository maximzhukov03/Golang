package main

import (
	"encoding/json"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/guptarohit/asciigraph"
	"github.com/gosuri/uilive"
	"net/http"
	"time"
)

type PriceResponse struct {
	Coin  string  `json:"symbol"`
	Price float64 `json:"price,string"`
}

var menu = map[string]string{
	"1": "BTCUSDT",
	"2": "LTCUSDT",
	"3": "ETHUSDT",
}

func getPrice(coin string) (float64, error) { // Получение курса монеты
	var priceResponse PriceResponse
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", coin)
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

func graf(symbol string, liveWriter *uilive.Writer) {
	data := []float64{}
	for {
		price, err := getPrice(symbol)
		if err != nil {
			price = 0
		}

		data = append(data, price)

		fmt.Print("\033[H\033[2J")
		graph := asciigraph.Plot(data, asciigraph.Width(100), asciigraph.Height(10))
		liveWriter.Write([]byte(graph)) // Записать новый график
		liveWriter.Flush() // Обновить экран

		time.Sleep(1 * time.Second)
	}
}

func getMenu() {
	fmt.Println("1. BTC_USD")
	fmt.Println("2. LTC_USD")
	fmt.Println("3. ETH_USD")
	fmt.Println("\nPress 1-3 to change symbol, press q to exit")
}

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	liveWriter := uilive.New()
	liveWriter.Start()

	currentSymbol := "BTCUSDT"

	for {
		getMenu()

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyBackspace {
			// Возвращаемся в меню
			continue
		}

		if key == keyboard.KeyEsc || char == 'q' {
			fmt.Println("\nВыход из программы...")
			return
		}

		if symbol, ok := menu[string(char)]; ok {
			currentSymbol = symbol
			// Стартуем новую горутину для отображения графика
			go graf(currentSymbol, liveWriter)
		}
	}
}