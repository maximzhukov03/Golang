package main

import (
	"fmt"
	"log"
	"time"
)

type GeneralIndicatorer interface {
	GetData(pair string, period int, from, to time.Time, indicator Indicatorer) ([]float64, error)
}

type Indicatorer interface {
	GetData(pair string, period int, from, to time.Time) ([]float64, error)
}

type GeneralIndicator struct{
	GeneralIndicatorer
}

func (genInd GeneralIndicator) GetData(pair string, period int, from, to time.Time, indicator Indicatorer) ([]float64, error){
	ema, err := indicator.GetData(pair, period, from, to)
	return ema, err
}

type SMA struct{
	SMAex Exchanger
}

func NewIndicatorSMA(ex Exchanger) *SMA{
	return &SMA{SMAex: ex}
}

func (s SMA) GetData(pair string, period int, from, to time.Time) ([]float64, error){
	closePrices, err := s.SMAex.GetClosePrice(pair, period, from, to)
	if err != nil {
		return nil, err
	}
	return calculateSMA(closePrices, 3), nil
}

type EMA struct{
	EMAex Exchanger
}

func NewIndicatorEMA(ex Exchanger) *EMA{
	return &EMA{EMAex: ex}
}

func (e EMA) GetData(pair string, period int, from, to time.Time) ([]float64, error){
	closePrices, err := e.EMAex.GetClosePrice(pair, period, from, to)
	if err != nil {
		return nil, err
	}
	s := calculateSMA(closePrices, 3)
	return calculateEMA(s, 3), nil
}


// Функция для расчета простого скользящего среднего (SMA)
func calculateSMA(data []float64, period int) []float64 {
	var sma = make([]float64, len(data)/period)
	for i := range sma {
		sum := 0.0
		for _, d := range data[i*period : i*period+period] {
			sum += d
		}
		sma[i] = sum / float64(period)
	}

	return sma
}

// Функция для расчета экспоненциального скользящего среднего (EMA)
func calculateEMA(data []float64, period int) []float64 {
	if len(data) == 0 || period <= 0 {
		return nil
	}

	alpha := 2.0 / (float64(period) + 1.0)
	ema := make([]float64, len(data))

	ema[0] = data[0]

	for i := 1; i < len(data); i++ {
		ema[i] = alpha*data[i] + (1-alpha)*ema[i-1]
	}

	return ema
}

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, resolution int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, resolution int, start, end time.Time) ([]float64, error)
}

func main() {
	var exchange Exchanger
	exchange = NewExmo()
	indicatorSMA := NewIndicatorSMA(exchange)
	generalIndicator := &GeneralIndicator{}
	sma, err := generalIndicator.GetData("BTC_USD", 30, time.Now().Add(-time.Hour*24*5), time.Now(), indicatorSMA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sma)

	indicatorEMA := NewIndicatorEMA(exchange)
	ema, err := generalIndicator.GetData("BTC_USD", 30, time.Now().Add(-time.Hour*24*5), time.Now(), indicatorEMA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ema)
}