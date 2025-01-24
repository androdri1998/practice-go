package main

import "fmt"

func times(multiplier int) func(value int) int {
	return func(value int) int {
		return multiplier * value
	}
}

func main() {
	multiplier := 2
	fmt.Println(multiplier)

	multiplierBy4 := times(4)
	multiplierBy5 := times(5)

	fmt.Println(multiplierBy4(4))
	fmt.Println(multiplierBy5(4))

}
