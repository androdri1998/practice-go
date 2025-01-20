package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	// slice it's not an array, slice it's a slice of an array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // new array but points to original array
	fmt.Println(a2, s3)

	// it's a structure that contains a size and a pointer to an element into an array
	s4 := s2[:1]
	fmt.Println(s2, s4)

}
