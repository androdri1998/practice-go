package main

import "fmt"

func main() {
	var a int
	var b float64
	var c int64
	var d bool
	var e rune
	var f byte
	var g *int
	var h string

	fmt.Printf("%v %v %v %v %v %v %v %q\n", a, b, c, d, e, f, g, h)
}
