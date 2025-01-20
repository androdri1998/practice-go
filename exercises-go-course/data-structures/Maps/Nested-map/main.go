package main

import "fmt"

func main() {
	employeesByLetter := map[string]map[string]float64{
		"J": {
			"Joe": 1000,
		},
		"D": {
			"Doe": 2000,
		},
	}

	fmt.Println(employeesByLetter)

	delete(employeesByLetter, "D")

	for letter, emploees := range employeesByLetter {
		fmt.Println(letter, emploees)
	}
}
