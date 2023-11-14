package main

import "fmt"

const a = 10
const b string = "Hello World"

func main() {
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
}
