package main

import "fmt"

type test_type int

var x test_type
var y int

func main() {
	fmt.Printf("%v %T\n", x, x)

	x = 42
	y = int(x)

	fmt.Println(x)
	fmt.Printf("%v %T\n", y, y)
}
