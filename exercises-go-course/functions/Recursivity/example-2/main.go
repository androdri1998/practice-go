package main

import "fmt"

func factorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		previousFactorial := factorial(n - 1)
		return n * previousFactorial
	}
}

func main() {
	fmt.Println(factorial(5))
}
