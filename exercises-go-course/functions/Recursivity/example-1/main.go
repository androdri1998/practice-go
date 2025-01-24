package main

import "fmt"

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Invalid number: %d", n)
	case n == 0:
		return 1, nil
	default:
		previousFactorial, _ := factorial(n - 1)
		return n * previousFactorial, nil
	}
}

func main() {
	fmt.Println(factorial(5))
	_, err := factorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
