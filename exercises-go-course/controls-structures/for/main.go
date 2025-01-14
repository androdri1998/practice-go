package main

import (
	"fmt"
	"time"
)

func main() {
	// for like while
	index := 0
	for index <= 10 {
		println(index)
		index++
	}

	fmt.Println("---------------------")
	// commom for
	for index = 0; index <= 20; index += 2 {
		fmt.Println(index)
	}

	fmt.Println("---------------------")
	for index = 1; index < 10; index++ {
		if index%2 == 0 {
			fmt.Println(index, "Even")
			continue
		}

		fmt.Println(index, "Odd")
	}

	fmt.Println("---------------------")
	// endless for
	index = 0
	for {
		fmt.Println("Forever", index)
		time.Sleep(time.Second)
		index++

		if index == 10 {
			break
		}
	}
}
