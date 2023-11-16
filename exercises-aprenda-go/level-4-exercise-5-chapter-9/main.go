package main

import "fmt"

func main() {
	list := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(list)

	list = append(list[:3], list[6:]...)

	fmt.Println(list)
}
