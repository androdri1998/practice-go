package main

import "fmt"

func main() {
	list := make([]int, 10, 10)

	for index := 1; index <= 10; index++ {
		list[index-1] = index * 2
	}

	fmt.Println(list[0:3])
	fmt.Println(list[4:])
	fmt.Println(list[1:8])
	fmt.Println(list[2 : len(list)-1])
}
