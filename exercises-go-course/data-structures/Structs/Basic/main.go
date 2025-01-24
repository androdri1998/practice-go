package main

import "fmt"

type Product struct {
	price    float64
	name     string
	discount float64
}

func (p Product) priceWithDiscountApplied() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	product1 := Product{20, "Product One", 0.3}
	product2 := Product{price: 30, name: "Product two", discount: 0.2}

	fmt.Println(product1)
	fmt.Printf("%.2f\n", product1.priceWithDiscountApplied())

	fmt.Println(product2)
	fmt.Printf("%.2f\n", product2.priceWithDiscountApplied())

}
