package main

import h "github.com/androdri1998/practice-go/exercises-go-course/fundamentals/functions/helpers"

func main() {
	result, err := h.Sum(2, 3)

	h.PrintInt(result)
	h.PrintError(err)
}
