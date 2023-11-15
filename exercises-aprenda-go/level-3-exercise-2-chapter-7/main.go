package main

import "fmt"

func main() {
	for index := 65; index <= 90; index++ {
		fmt.Printf("%v ------\n", index)

		for intenal_index := 0; intenal_index < 3; intenal_index++ {
			fmt.Printf("%#U\n", index)
		}
	}
}
