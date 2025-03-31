package main

import (
	"fmt"
	"errors"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

type DineInOrder struct{
	orderDetails map[string]int
}

type TakeAwayOrder struct{
	orderDetails map[string]int
}

func (d *DineInOrder) GetOrderDetails() map[string]int{
	return d.orderDetails
}

func (d *DineInOrder) AddItem(food string, count int) error{ 
	if count > 0{
		d.orderDetails[food] = count
		d.GetOrderDetails()
		return nil
		}else{
		return errors.New("")
	}
}

func (d *DineInOrder) RemoveItem(food string) error{
	_, ok := d.orderDetails[food]
	if ok{
		delete(d.orderDetails, food)	
		d.GetOrderDetails()
		return nil	
	}else{
		return errors.New("")
	}
}

func (d *TakeAwayOrder) GetOrderDetails() map[string]int{
	return d.orderDetails
}

func (d *TakeAwayOrder) AddItem(food string, count int) error{ 
	if count > 0{
		d.orderDetails[food] = count
		d.GetOrderDetails()
		return nil
		}else{
		return errors.New("")
	}
}

func (d *TakeAwayOrder) RemoveItem(food string) error{
	_, ok := d.orderDetails[food]
	if ok{
		delete(d.orderDetails, food)	
		d.GetOrderDetails()
		return nil	
	}else{
		return errors.New("")
	}
}

func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	ManageOrder(dineIn)
	ManageOrder(takeAway)
}