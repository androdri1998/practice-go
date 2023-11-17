package main

import "fmt"

type person struct {
	name     string
	lastname string
	flavors  []string
}

func showPerson(people map[string]person) {
	for _, uniquePerson := range people {
		fmt.Println(uniquePerson.name)
		fmt.Println(uniquePerson.lastname)

		for _, flavor := range uniquePerson.flavors {
			fmt.Println(flavor)
		}

		fmt.Println("---------")
	}
}

func main() {
	people := map[string]person{
		"lastname1": {name: "firstname 1",
			lastname: "lastname 1",
			flavors:  []string{"flavor 1", "flavor 2"},
		},
		"lastname2": {
			name:     "firstname 2",
			lastname: "lastname 2",
			flavors:  []string{"flavor 3", "flavor 4"},
		},
	}
	showPerson(people)
}
