package main

import "fmt"

func main() {
	x := 9

	switch {
	case x == 10:
		fmt.Printf("X it's equal to 10\n")
	case x == 9:
		fmt.Printf("X it's equal to 9\n")
	default:
		fmt.Printf("X it's unknown")
	}
}
