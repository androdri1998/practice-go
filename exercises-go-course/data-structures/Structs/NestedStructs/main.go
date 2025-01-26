package main

import "fmt"

type item struct {
	productId int
	amount    int
	price     float64
}

type order struct {
	userId int
	items  []item
}

func (o order) totalValue() float64 {
	total := 0.0
	for _, item := range o.items {
		total += item.price * float64(item.amount)
	}

	return total
}

func main() {
	order := order{
		userId: 1,
		items:  []item{item{1, 2, 12.10}, item{2, 1, 23.49}, item{11, 100, 3.13}},
	}

	fmt.Println(order.totalValue())
}
