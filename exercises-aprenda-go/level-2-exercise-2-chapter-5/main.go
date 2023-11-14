package main

import "fmt"

func main() {
	equal := 10 == 10
	smallerOrEqual := 9 <= 8
	greaterOrEqual := 8 >= 9
	diff := 1 != 2
	smaller := 9 < 10
	greater := 10 > 9

	fmt.Printf("%v %T\n", equal, equal)
	fmt.Printf("%v %T\n", smallerOrEqual, smallerOrEqual)
	fmt.Printf("%v %T\n", greaterOrEqual, greaterOrEqual)
	fmt.Printf("%v %T\n", diff, diff)
	fmt.Printf("%v %T\n", smaller, smaller)
	fmt.Printf("%v %T\n", greater, greater)
}
