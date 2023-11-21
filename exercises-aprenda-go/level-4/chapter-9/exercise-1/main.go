package main

import "fmt"

func main() {
	list := [5]int{1, 2, 3, 4, 5}

	for _, value := range list {
		fmt.Println(value)
	}

	fmt.Printf("%T\n", list)
}
