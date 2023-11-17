package main

import "fmt"

type person struct {
	name     string
	lastname string
	flavors  []string
}

func showPerson(people person) {
	fmt.Println(people.name)
	fmt.Println(people.lastname)

	for _, flavor := range people.flavors {
		fmt.Println(flavor)
	}

	fmt.Println("---------")
}

func main() {
	person1 := person{
		name:     "firstname 1",
		lastname: "lastname 1",
		flavors:  []string{"flavor 1", "flavor 2"},
	}

	person2 := person{
		name:     "firstname 2",
		lastname: "lastname 2",
		flavors:  []string{"flavor 3", "flavor 4"},
	}

	showPerson(person1)
	showPerson(person2)
}
