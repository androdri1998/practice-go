package main

import "fmt"

type test_type int

var x test_type

func main() {
	fmt.Printf("%v %T\n", x, x)

	x = 42
	fmt.Println(x)
}
