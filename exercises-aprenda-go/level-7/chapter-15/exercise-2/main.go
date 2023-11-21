package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	me := person{"my name", 21}

	fmt.Println(me)
	changeMe(&me)
	fmt.Println(me)
}

func changeMe(someone *person) {
	(*someone).age = 22
	someone.name = "another name"
}
