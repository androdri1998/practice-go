package main

import "fmt"

func main() {
	employeesAndPaycheck := map[string]float64{
		"Joe": 1000.25,
		"Doe": 2000.50,
	}

	employeesAndPaycheck["Other Joe"] = 1320

	delete(employeesAndPaycheck, "non-existent")

	fmt.Println(employeesAndPaycheck)

	for name, paycheck := range employeesAndPaycheck {
		fmt.Println(name, paycheck)
	}
}
