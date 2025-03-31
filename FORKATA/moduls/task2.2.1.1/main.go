package main

import (
	"errors"
	"fmt"
)

type Accounter interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type CurrentAccount struct {
	balance float64
}

type SavingsAccount struct{
	balance float64
}


func (a *CurrentAccount) Deposit(money float64) error{
	if money > 0{
		a.balance = a.balance + float64(money)
		return nil
	}else{
		return errors.New("Сумма депозита введена неверно")
	}
} 

func (a *CurrentAccount) Withdraw(money float64) error{
	if a.balance < 500{
		return errors.New("Нельзя снимать если на балансе меньше 500")
	}else if a.balance - money < 0{
		return errors.New("Нельзя снимать сумму больше чем лежит на балансе")
	}else{
		a.balance = a.balance - float64(money)
		return nil
	}
} 

func (a *CurrentAccount) Balance() (float64){
	return a.balance
}

func (a *SavingsAccount) Deposit(money float64) error{
	if money > 0{
		a.balance = a.balance + float64(money)
		return nil
	}else{
		return errors.New("Сумма депозита введена неверно")
	}
} 

func (a *SavingsAccount) Withdraw(money float64) error{
	if a.balance < 500{
		return errors.New("Нельзя снимать если на балансе меньше 500")
	}else if a.balance - money < 0{
		return errors.New("Нельзя снимать сумму больше чем лежит на балансе")
	}else{
		a.balance = a.balance - float64(money)
		return nil
	}
} 

func (a *SavingsAccount) Balance() (float64){
	return a.balance
} 

func ProcessAccount(a Accounter) {
	a.Deposit(700)
	a.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", a.Balance())
}

func main() {
	c := &CurrentAccount{}
	s := &SavingsAccount{}
	ProcessAccount(c)
	ProcessAccount(s)
}