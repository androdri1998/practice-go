package main

import "fmt"

func main() {
	for index := 33; index <= 122; index++ {
		fmt.Printf("%v\t %d\t %#x\t %#U\n", string(rune(index)), index, index, index)
	}
}
