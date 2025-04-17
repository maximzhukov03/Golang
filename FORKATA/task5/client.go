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

type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo(opts ...func(exmo *Exmo)) Exchanger {
	exmo := &Exmo{
		client: http.DefaultClient,
		url:    "https://api.exmo.com/v1",
	}
	for _, opt := range opts {
		opt(exmo)
	}
	
	return exmo
}

func GetConv(constanta string, url url.Values, exmo *Exmo) ([]byte, error) {
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

func (e *Exmo) GetTicker() (Ticker, error) {
	var tick Ticker

	body, err := GetConv(ticker, nil, e)

	err = json.Unmarshal(body, &tick)
	if err != nil {
		return tick, err
	}
	return tick, err
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
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

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
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

func (e *Exmo) GetCurrencies() (Currencies, error) {
	var curr Currencies
	body, err := GetConv(currency, nil, e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &curr)
	if err != nil {
		return curr, err
	}
	return curr, err
}

func (e *Exmo) GetCandlesHistory(pair string, resolution int, from, to time.Time) (CandlesHistory, error) {
	var candles CandlesHistory

	params := url.Values{}
	params.Set("symbol", pair)
	params.Set("resolution", fmt.Sprintf("%d", resolution))
	params.Set("from", fmt.Sprintf("%d", from.Unix()))
	params.Set("to", fmt.Sprintf("%d", to.Unix()))

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

func (e *Exmo) GetClosePrice(pair string, resolution  int, start, end time.Time) ([]float64, error) {
	candles, err := e.GetCandlesHistory(pair, resolution , start, end)
	if err != nil {
		return nil, err
	}
	var closes []float64
	for _, c := range candles.Candles {
		closes = append(closes, c.C)
	}
	return closes, nil
}
