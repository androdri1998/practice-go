package main

import "fmt"

func main() {
	for index := 33; index <= 122; index++ {
		fmt.Printf("%d\t %#x\t %#U\n", index, index, index)
	}
}
