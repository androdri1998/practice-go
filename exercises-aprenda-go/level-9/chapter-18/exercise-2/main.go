package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Printf("my name is %s\n", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	person1 := person{
		name: "Some name",
	}

	person1.speak()        // this is a shortcut for (&person1).speak()
	(&person1).speak()     // this is the way to call through a pointer
	saySomething(&person1) // work
	// saySomething(person1)  // does not work
}
