package main

import "fmt"

func main() {
	for index := 0; index <= 10; index++ {

		result := index % 2
		if result == 0 {
			fmt.Printf("There's no remain value for %v\n", index)
			continue
		}

		fmt.Printf("There's %v remain value for %v\n", result, index)
	}
}
