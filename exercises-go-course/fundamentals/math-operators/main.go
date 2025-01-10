package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 3, 2
	fmt.Println("Sum:", a+b)
	fmt.Println("Sub:", a-b)
	fmt.Println("Div:", a/b)
	fmt.Println("Times:", a*b)
	fmt.Println("Mod:", a%b)

	// bitwise
	fmt.Println("E:", a&b)
	fmt.Println("OR:", a|b)
	fmt.Println("XOR:", a^b)

	c := 3.0
	d := 2.0
	fmt.Println("Bigger:", math.Max(float64(a), float64(b)))
	fmt.Println("Menor:", math.Min(c, d))
	fmt.Println("Bigger:", math.Pow(c, d))
}
