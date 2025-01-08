package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("integer literal is:", reflect.TypeOf(32000))

	// types to integers only positive numbers: uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("byte literal is:", reflect.TypeOf(b))

	// types to integers numbers: int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("max int it's:", i1)
	fmt.Println("type of i1 is:", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("type of rune is:", reflect.TypeOf(i2))
	fmt.Println("rune is:", i2)

	// real numbers float32, float64
	var x float32 = 40.12
	y := float32(40.12)
	fmt.Println("type of x is:", reflect.TypeOf(x))
	fmt.Println("type of y is:", reflect.TypeOf(y))
	fmt.Println("type of 49.99 is:", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("type of bo is", reflect.TypeOf(bo))
	fmt.Println("negative of bo is", !bo)

	// string
	s1 := "this is a string"
	fmt.Println(s1 + " end")
	fmt.Println("length of string", len(s1))

	// strings multiline
	s2 := `this
	is
		another
		string`
	fmt.Println(s2)

	// single letter
	char := 'a'
	fmt.Println(char)
	fmt.Println(reflect.TypeOf(char))
}
