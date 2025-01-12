package main

import "fmt"

func main() {
	i := 1

	// Go lang there's not arithmetic of pointers
	var iPointer *int = nil
	fmt.Println(iPointer)

	iPointer = &i

	fmt.Println(i, iPointer)

	*iPointer++
	i++
	fmt.Println(iPointer, *iPointer, i, &i)
}
