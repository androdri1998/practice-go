package main

import "fmt"

func main() {
	fmt.Print("First")
	fmt.Print(" line\n")

	fmt.Println("New")
	fmt.Println(" line")

	x := 3.1415

	fmt.Println("X value it's", x, "other things")

	xs := fmt.Sprint(x)

	fmt.Println("X value it's " + xs)

	fmt.Printf("X value it's %.2f\n", x)

	a, b, c, d := 1, 1.9999, false, "something"

	fmt.Printf("\n%d %f %.1f %t %s\n", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
