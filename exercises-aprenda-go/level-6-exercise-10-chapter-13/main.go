package main

import "fmt"

func main() {
	incrementBy2 := incrementBy(2)

	fmt.Println(incrementBy2())
	fmt.Println(incrementBy2())
	fmt.Println(incrementBy2())
	fmt.Println(incrementBy2())
}

func incrementBy(number int) func() int {
	incrementValue := number
	total := 0 - number

	return func() int {
		total += incrementValue

		return total
	}
}
