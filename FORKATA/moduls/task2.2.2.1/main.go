package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func NewOrder(id int, options ...OrderOption) *Order {
    o := &Order{
        ID: id,
    }

    for _, option := range options {
        option(o)
    }

    return o
}

func WithCustomerID(id string) OrderOption {
    return func(p *Order) {
        p.CustomerID = id
    }
}

func WithItems(item []string) OrderOption {
    return func(p *Order) {
        p.Items = item
    }
}

func WithOrderDate(time time.Time) OrderOption {
    return func(p *Order) {
        p.OrderDate = time
    }
}

func main() {
	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(time.Now()))

	fmt.Printf("Order: %+v\n", order)
}