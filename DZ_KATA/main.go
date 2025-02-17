package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"os"
)

func main() {
	fakeGen := GenerateFakeData()
	fmt.Println(fakeGen)
	os.Exit(0)
}

func GenerateFakeData() string {
	Name := fake.FullName()
	Address := fake.StreetAddress()
	Phone := fake.Phone()
	Email := fake.EmailAddress()
	fakeGen := fmt.Sprintf("Name: %s\nAddress: %s\nPhone: %s\nEmail: %s", Name, Address, Phone, Email)
	return fakeGen
}