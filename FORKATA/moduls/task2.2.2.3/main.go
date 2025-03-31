package main

import (
	"fmt"
	"sync"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

type Customer struct {
	ID       int
	Name     string
	Account  Account
}

type CheckingAccount struct {
	balance float64
	mu      sync.Mutex
}

type SavingsAccount struct {
	balance float64
	mu      sync.Mutex
}

func (s *SavingsAccount) Deposit(money float64){
	s.mu.Lock()
	defer s.mu.Unlock()
	s.balance += money
	
}

func (s *SavingsAccount) Withdraw(money float64) error{
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.balance < 1000 && (s.balance - money) >= 0{
		return fmt.Errorf("Баланс меньше 1000")
	}else{
		s.balance -= money
		return nil
	}
	
}

func (s *SavingsAccount) Balance() float64{
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.balance
}

func (s *CheckingAccount) Deposit(money float64){
	s.mu.Lock()
	defer s.mu.Unlock()
	s.balance += money
}

func (s *CheckingAccount) Withdraw(money float64) error{
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.balance < money && (s.balance - money) >= 0{
		return fmt.Errorf("Баланс меньше 1000")
	}else{
		s.balance -= money
		return nil
	}
	
}

func (s *CheckingAccount) Balance(){
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println(s.balance) 
}

type CustomerOption func(*Customer)

func NewCustomer(id int, options ...CustomerOption) *Customer {
	customer := &Customer{
		ID: id,
	}

	for _, option := range options {
		option(customer)
	}

	return customer
}


func WithName(name string) CustomerOption{
	return func(c *Customer){
		c.Name = name
	}
}

func WithAccount(account Account) CustomerOption {
	return func(c *Customer) {
		c.Account = account
	}
}

func main() {
	savings := &SavingsAccount{}
	savings.Deposit(1000)

	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))
	
	err := customer.Account.Withdraw(100)
	if err != nil {
        fmt.Println(err)
    }

	fmt.Printf("Customer: %v, Balance: %v\n", customer.Name, customer.Account.Balance())
}