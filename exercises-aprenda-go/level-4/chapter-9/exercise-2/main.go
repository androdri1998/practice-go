package main

import "fmt"

func main() {
	list := make([]int, 10, 10)

	for index := 1; index <= 10; index++ {
		list[index-1] = index * 2
	}

	for _, value := range list {
		fmt.Println(value)
	}

	fmt.Printf("%T\n", list)
}
