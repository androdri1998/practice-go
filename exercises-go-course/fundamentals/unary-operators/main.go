package main

import "fmt"

func main() {
	x := 1
	y := 2

	x++
	fmt.Println("x:", x)
	x++
	fmt.Println("x:", x)
	x++
	fmt.Println("x:", x)

	y--
	fmt.Println("y:", y)
	y--
	fmt.Println("y:", y)
}
