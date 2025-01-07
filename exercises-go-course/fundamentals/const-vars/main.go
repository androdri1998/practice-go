package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var ray = 3.2

	area := m.Pow(ray, 2) * PI

	fmt.Println("Area is", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "something"

	fmt.Println(g, h, i)
}
