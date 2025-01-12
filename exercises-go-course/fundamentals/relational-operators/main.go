package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Dates:", d1 == d2)
	fmt.Println("Dates:", d1.Equal(d2))

	type Person struct {
		Nome string
	}

	p1 := Person{"Joe"}
	p2 := Person{"Other Joe"}
	p3 := Person{"Joe"}

	fmt.Println("Structs:", p1 == p3)
	fmt.Println("Structs:", p1 == p2)
}
