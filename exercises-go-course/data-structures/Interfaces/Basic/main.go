package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastname string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return p.name + " " + p.lastname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - %.2f", p.name, p.price)
}

func print(file printable) {
	fmt.Println(file.toString())
}

func main() {
	var something printable = person{"Joe", "Doe"}
	print(something)

	something = product{"Product 1", 20.0}
	print(something)
}
