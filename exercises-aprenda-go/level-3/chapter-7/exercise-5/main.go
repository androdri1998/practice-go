package main

import "fmt"

func main() {
	for index := 10; index <= 100; index++ {
		result := index % 4

		fmt.Printf("%v - %v\n", index, result)
	}
}
