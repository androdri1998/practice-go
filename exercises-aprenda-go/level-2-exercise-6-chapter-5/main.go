package main

import "fmt"

const (
	_           = iota
	first_year  = iota
	second_year = iota
	third_year  = iota
	fourth_year = iota
)

func main() {
	fmt.Println(first_year)
	fmt.Println(second_year)
	fmt.Println(third_year)
	fmt.Println(fourth_year)
}
