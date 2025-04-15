package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/cinar/indicator"
)

type Indicator interface {
	StochPrice() ([]float64, []float64)
	RSI(period int) ([]float64, []float64)
	StochRSI(rsiPeriod int) ([]float64, []float64)
	SMA(period int) []float64
	MACD() ([]float64, []float64)
	EMA() []float64
}

func UnmarshalKLines(data []byte) (KLines, error) {
	var r KLines
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *KLines) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type KLines struct {
	Pair    string   `json:"pair"`
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

type Lines struct {
	high    []float64
	low     []float64
	closing []float64
}

func (t *Lines) StochPrice() ([]float64, []float64) {
	k, d := indicator.StochasticOscillator(t.high, t.low, t.closing)
	return k, d
}

func (t *Lines) RSI(period int) ([]float64, []float64) {
	rs, rsi := indicator.RsiPeriod(period, t.closing)

	return rs, rsi
}

func (t *Lines) StochRSI(rsiPeriod int) ([]float64, []float64) {
	_, rsi := t.RSI(rsiPeriod)
	k, d := indicator.StochasticOscillator(rsi, rsi, rsi)

	return k, d
}

func (t *Lines) SMA(period int) []float64 {

	return indicator.Sma(period, t.closing)
}

func (t *Lines) MACD() ([]float64, []float64) {
	return indicator.Macd(t.closing)
}

func (t *Lines) EMA() []float64 {
	return indicator.Ema(5, t.closing)
}

type LinesProxy struct {
	lines Indicator
	cache map[string][]float64
}

func LoadKlinesProxy(data []byte) *LinesProxy {
	klines := LoadKlines(data)

	lineProx := &LinesProxy{
		lines: klines,
		cache: make(map[string][]float64),
	}
	return lineProx
}

func LoadKlines(data []byte) *Lines {
	klines, err := UnmarshalKLines(data)
	if err != nil {
		log.Fatal(err)
	}

	t := &Lines{}
	for _, v := range klines.Candles {
		t.closing = append(t.closing, v.C)
		t.low = append(t.low, v.L)
		t.high = append(t.high, v.H)
	}

	return t
}

func (lP *LinesProxy) StochPrice() ([]float64, []float64) {
	key1 := "k_stochprice"
	key2 := "d_stochprice"
	
	k, ok := lP.cache[key1]
	r, ok2 := lP.cache[key2]
	if ok && ok2{
		return k, r
	}

	k, r = lP.lines.StochPrice()
	lP.cache[key1] = k
	lP.cache[key2] = r
	return k, r

}

func (lP *LinesProxy) RSI(period int) ([]float64, []float64) {
	key1 := fmt.Sprintf("rs_%v", period)
	key2 := fmt.Sprintf("rsi_%v", period)
	
	rs, ok := lP.cache[key1]
	rsi, ok2 := lP.cache[key2]
	if ok && ok2{
		return rs, rsi
	}

	rs, rsi = lP.lines.RSI(period)
	lP.cache[key1] = rs
	 lP.cache[key2] = rsi
	return rs, rsi
}

func (lP *LinesProxy) StochRSI(rsiPeriod int) ([]float64, []float64) {
	key1 := fmt.Sprintf("k_stochrsi_%v", rsiPeriod)
	key2 := fmt.Sprintf("d_stochrsi_%v", rsiPeriod)
	
	krsi, ok := lP.cache[key1]
	drsi, ok2 := lP.cache[key2]
	if ok && ok2{
		return krsi, drsi
	}
	krsi, drsi = lP.lines.StochRSI(rsiPeriod)
	lP.cache[key1] = krsi
	lP.cache[key2] = drsi
	return krsi, drsi
}

func (lP *LinesProxy) SMA(period int) []float64 {
	key1 := fmt.Sprintf("sma_%v", period)
	
	sma, ok := lP.cache[key1]
	if !ok{
		sma = lP.lines.SMA(period)
		lP.cache[key1] = sma
		return sma
	}
	return sma
}

func (lP *LinesProxy) MACD() ([]float64, []float64) {
	key1 := fmt.Sprintf("macd")
	key2 := fmt.Sprintf("signal")
	
	macd, ok := lP.cache[key1]
	signal, ok2 := lP.cache[key2]
	if ok && ok2{
		return macd, signal
	}
	macd, signal = lP.lines.MACD()
	lP.cache[key1] = macd
	lP.cache[key2] = signal
	return macd, signal
}

func (lP *LinesProxy) EMA() []float64 {
	key1 := fmt.Sprintf("ema")
	
	ema, ok := lP.cache[key1]
	if !ok{
		ema = lP.lines.EMA()
		lP.cache[key1] = ema
		return ema
	}
	return ema
}

func LoadCandles(pair string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.exmo.com/v1.1/candles_history?symbol=%s&resolution=30&from=1703056979&to=1705476839", pair), nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	pair := "BTC_USD"
	candles := LoadCandles(pair)
	lines := LoadKlinesProxy(candles)
	lines.RSI(3)
}