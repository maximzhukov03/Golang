package main

import (
	"fmt"
	"time"
	"github.com/cinar/indicator"
)

type Indicatorer interface {
	SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
	EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
}

type IndicatorOption func(*Indicator)

type Indicator struct {
	exchange     Exchanger
	calculateSMA func(data []float64, period int) []float64
	calculateEMA func(data []float64, period int) []float64
}

func WithCalculateEMA(e func([]float64, int) []float64) IndicatorOption{
	return func(i *Indicator){
		i.calculateEMA = e
	}
}

func WithCalculateSMA(e func([]float64, int) []float64) IndicatorOption{
	return func(i *Indicator){
		i.calculateSMA = e
	}
}


func calculateSMA(closing []float64, period int) []float64 {
	return indicator.Sma(period, closing)
}

func calculateEMA(closing []float64, period int) []float64 {
	return indicator.Ema(period, closing)
}

func (i *Indicator) SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
    prices, err := i.exchange.GetClosePrice(pair, resolution, from, to)
    if err != nil {
        return nil, err
    }

    res := i.calculateSMA(prices, period)
    
    return res, nil
}

func (i *Indicator) EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
    prices, err := i.exchange.GetClosePrice(pair, resolution, from, to)
    if err != nil {
        return nil, err
    }

    res := i.calculateEMA(prices, period)
    
    return res, nil
}

func NewIndicator(exchange Exchanger, opts ...IndicatorOption) Indicatorer{
	ind := &Indicator{
		exchange: exchange,
	}

	for _, opt := range opts{
		opt(ind)
	}
	return ind
}

func main() {
	var exchange Exchanger
	exchange = NewExmo()
	ind := NewIndicator(exchange, WithCalculateEMA(calculateEMA), WithCalculateSMA(calculateSMA))
	sma, err := ind.SMA("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sma)
	ema, err := ind.EMA("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ema)
}