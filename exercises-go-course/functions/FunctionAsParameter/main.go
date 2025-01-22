package main

import "fmt"

func times(a, b int) int {
	return a * b
}

func exec(function func(a, b int) int, p1, p2 int) int {
	result := function(p1, p2)
	return result
}

func main() {
	fmt.Println(exec(times, 2, 4))
}
