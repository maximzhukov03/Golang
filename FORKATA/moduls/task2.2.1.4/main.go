package main

import (
	"fmt"
	"errors"
)

type PaymentMethod interface {
	Pay(amount float64) error
}

type CreditCard struct{
	balance float64
}

type Bitcoin struct{
	balance float64
}

func (p *CreditCard) Pay(a float64) error{
	if p.balance < a{
		return errors.New("недостаточный баланс")
	}else if a < 0{
		return errors.New("недопустимая сумма платежа")
	}else{
		p.balance = p.balance - a
		fmt.Printf("Оплачено %g с помощью кредитной карты\n", a)
		return nil
	}
}

func (p *Bitcoin) Pay(a float64) error{
	if p.balance < a{
		return errors.New("недостаточный баланс")
	}else if a < 0{
		return errors.New("недопустимая сумма платежа")
	}else{
		p.balance = p.balance - a
		fmt.Printf("Оплачено %g с помощью биткоина\n", a)
		return nil
	}
}

func ProcessPayment(p PaymentMethod, amount float64) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println("Не удалось обработать платеж:", err)
	}
}

func main() {
	cc := &CreditCard{balance: 500.00}
	btc := &Bitcoin{balance: 2.00}

	ProcessPayment(cc,  200.00)
	ProcessPayment(btc, 1.00)
}