package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	fmt.Println("Test " + string(97))

	fmt.Println("Test " + strconv.Itoa(123))

	num, err := strconv.Atoi("123")
	fmt.Println(num-122, err)

	b, err := strconv.ParseBool(("true"))
	if b {
		fmt.Println(("is true"))
	}
}
