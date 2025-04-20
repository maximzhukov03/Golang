package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"sort"
	"time"
)

type Product struct {
	Name      string
	Price     float64
	CreatedAt time.Time
	Count     int
}

func (p Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %f, Count: %v", p.Name, p.Price, p.Count)
}

func generateProducts(n int) []Product {
	gofakeit.Seed(time.Now().UnixNano())
	products := make([]Product, n)
	for i := range products {
		products[i] = Product{
			Name:      gofakeit.Word(),
			Price:     gofakeit.Price(1.0, 100.0),
			CreatedAt: gofakeit.Date(),
			Count:     gofakeit.Number(1, 100),
		}
	}
	return products
}

type ByPrice []Product
type ByCount []Product
type ByCreatedAt []Product


func (bp ByPrice) Len() int { 
	return len(bp) 
}
func (bp ByPrice) Less(i, j int) bool {
	return bp[i].Price < bp[j].Price 
}
func (bp ByPrice) Swap(i, j int) {
	bp[i], bp[j] = bp[j], bp[i]
}

func (bc ByCount) Len() int { 
	return len(bc) 
}
func (bc ByCount) Less(i, j int) bool {
	return bc[i].Count < bc[j].Count 
}
func (bc ByCount) Swap(i, j int) {
	bc[i], bc[j] = bc[j], bc[i]
}

func (bc ByCreatedAt) Len() int { 
	return len(bc) 
}
func (bc ByCreatedAt) Less(i, j int) bool {
	res := bc[i].CreatedAt.Before(bc[j].CreatedAt)
	return res
}
func (bc ByCreatedAt) Swap(i, j int) {
	bc[i], bc[j] = bc[j], bc[i]
}


func main() {
	products := generateProducts(10)

	fmt.Println("Исходный список:")
	fmt.Println(products)

	// Сортировка продуктов по цене
	sort.Sort(ByPrice(products))
	fmt.Println("\nОтсортировано по цене:")
	fmt.Println(products)

	// Сортировка продуктов по дате создания
	sort.Sort(ByCreatedAt(products))
	fmt.Println("\nОтсортировано по дате создания:")
	fmt.Println(products)

	// Сортировка продуктов по количеству
	sort.Sort(ByCount(products))
	fmt.Println("\nОтсортировано по количеству:")
	fmt.Println(products)
}