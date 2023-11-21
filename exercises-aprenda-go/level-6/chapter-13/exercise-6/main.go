package main

import "fmt"

func main() {
	func(number int) {
		fmt.Println(number * 2)
	}(4)
}
