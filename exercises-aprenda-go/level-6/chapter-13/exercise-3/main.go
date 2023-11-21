package main

import "fmt"

func main() {
	defer fmt.Println("start")
	for index := 0; index < 10; index++ {
		fmt.Println(index)
	}
	fmt.Println("end")
}
