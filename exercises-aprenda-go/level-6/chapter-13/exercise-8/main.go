package main

import "fmt"

func main() {
	multiplyBy2 := multiplyBy(2)

	fmt.Println(multiplyBy2(2))
}

func multiplyBy(number int) func(n int) int {
	return func(n int) int {
		return n * number
	}
}
