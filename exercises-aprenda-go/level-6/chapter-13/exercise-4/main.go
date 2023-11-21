package main

type person struct {
	name     string
	lastname string
	age      int
}

func (p person) info() {
	println(p.name, p.lastname, p.age)
}

func main() {
	myself := person{"my name", "my lastname", 21}

	myself.info()
}
