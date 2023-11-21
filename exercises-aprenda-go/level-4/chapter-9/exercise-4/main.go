package main

import "fmt"

func main() {
	list := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	list = append(list, 52)
	list = append(list, 53, 54, 55)

	fmt.Println(list)

	anotherSlice := []int{56, 57, 58, 59, 60}

	list = append(list, anotherSlice...)

	fmt.Println(list)
}
