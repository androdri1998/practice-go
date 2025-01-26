package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastname string
}

func (p person) getFullName() string {
	return p.name + " " + p.lastname
}

func (p *person) seFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.name = parts[0]
	p.lastname = parts[1]
}

func main() {
	p1 := person{"Joe", "Doe"}

	fmt.Println(p1.getFullName())

	p1.seFullName("otherName otherLastname")
	fmt.Println(p1.getFullName())
}
