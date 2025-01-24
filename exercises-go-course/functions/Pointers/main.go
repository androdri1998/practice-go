package main

import "fmt"

func start1(n int) {
	n++
}

func start2(n *int) {
	*n++
}

func main() {
	n := 0
	fmt.Println(n)

	start1(n)
	start2(&n)

	fmt.Println(n)
}
