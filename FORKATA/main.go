package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

type Ticker map[string]TickerValue // Структура для /ticker
type TickerValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

type Trades map[string][]Pair // Структура для /trades

type Pair struct {
	TradeID  int64  `json:"trade_id"`
	Date     int64  `json:"date"`
	Type     string `json:"type"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
}

type OrderBook map[string]OrderBookPair // Структура для /order_book

type OrderBookPair struct{
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BidQuantity string     `json:"bid_quantity"`
	BidAmount   string     `json:"bid_amount"`
	BidTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}

type Currencies []string

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}




type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(symbol string, resolution int, from, to int) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

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

func GetConv(constanta string, url url.Values, exmo *Exmo) ([]byte, error){
	client := exmo.client
	urlReq := exmo.url + constanta
	if url != nil {
		urlReq += "?" + url.Encode()
	}

	fmt.Println(urlReq)
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

func (e *Exmo) GetTicker() (Ticker, error) {
	var tick Ticker

	body, err := GetConv(ticker, nil, e)

	err = json.Unmarshal(body, &tick)
	if err != nil {
		return tick, err
	}
	return tick, err
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error){
	var trade Trades
	
	url := url.Values{}
	url.Set("pair", strings.Join(pairs, ","))
	
	body, err := GetConv(trades, url, e)

	err = json.Unmarshal(body, &trade)
	if err != nil {
		return trade, err
	}
	return trade, err
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error){
	var order OrderBook
	url := url.Values{}
	url.Set("pair", strings.Join(pairs, ","))
	url.Set("limit", fmt.Sprintf("%d", limit))
	body, err := GetConv(orderBook, url, e)

	err = json.Unmarshal(body, &order)
	if err != nil {
		return order, err
	}
	return order, err
}

func (e *Exmo) GetCurrencies() (Currencies, error){
	var curr Currencies
	body, err := GetConv(currency, nil, e)
	if err !=  nil{
		return nil, err
	}
	err = json.Unmarshal(body, &curr)
	if err != nil {
		return curr, err
	}
	return curr, err
}

func (e *Exmo) GetCandlesHistory(symbol string, resolution int, from, to any) (CandlesHistory, error) {
	var candles CandlesHistory

	var fromUnix, toUnix int

    switch v := from.(type) {
    case time.Time:
        fromUnix = int(v.Unix())
    case int:
        fromUnix = v
    default:
        return CandlesHistory{}, fmt.Errorf("unsupported 'from' type: %T", from)
    }

    switch v := to.(type) {
    case time.Time:
        toUnix = int(v.Unix())
    case int:
        toUnix = v
    default:
        return CandlesHistory{}, fmt.Errorf("unsupported 'to' type: %T", to)
    }

	params := url.Values{}
	params.Set("symbol", symbol)
	params.Set("resolution", fmt.Sprintf("%d", resolution))
	params.Set("from", fmt.Sprintf("%d", fromUnix))
	params.Set("to", fmt.Sprintf("%d", toUnix))

	body, err := GetConv(candlesHistory, params, e)
	if err != nil {
		return candles, err
	}

	err = json.Unmarshal(body, &candles)
	if err != nil {
		return candles, err
	}
	return candles, nil
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error){
	candles, err := e.GetCandlesHistory(pair, limit, int(start.Unix()), int(end.Unix()))
	if err != nil {
		return nil, err
	}
	var closes []float64
	for _, c := range candles.Candles {
		closes = append(closes, c.C)
	}
	return closes, nil
}




func main() {
	exchange := NewExmo()
	candles, err := exchange.GetCandlesHistory("BTC_USD", 15, int(time.Now().Add(-time.Hour*24).Unix()), int(time.Now().Unix()))
	if err != nil {
		return
	}
	jsonBytes, err := json.MarshalIndent(candles, "", "\t")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonBytes))
}