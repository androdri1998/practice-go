package main

import "fmt"

const (
	_ = 2023 + iota
	first_year
	second_year
	third_year
	fourth_year
)

func main() {
	fmt.Println(first_year)
	fmt.Println(second_year)
	fmt.Println(third_year)
	fmt.Println(fourth_year)
}
