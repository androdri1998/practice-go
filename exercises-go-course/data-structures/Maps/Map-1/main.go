package main

import "fmt"

func main() {
	// var approved map[int]string
	// maps should be initialized
	approved := make(map[int]string)

	approved[1] = "Joe"
	approved[2] = "Doe"

	fmt.Println(approved)

	for key, name := range approved {
		fmt.Println(key, name)
	}

	fmt.Println(approved[1])

	delete(approved, 1)

	fmt.Println(approved[1])
}
