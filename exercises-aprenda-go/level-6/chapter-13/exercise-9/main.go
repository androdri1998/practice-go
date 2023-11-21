package main

import "fmt"

func main() {
	fmt.Println(calc(sum, 2, 3))
}

func sum(a, b int) int {
	return a + b
}

func calc(calcFunc func(a, b int) int, a int, b int) int {
	result := calcFunc(a, b)

	return result
}
