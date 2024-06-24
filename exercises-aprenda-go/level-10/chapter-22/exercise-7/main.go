package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)

	for indexExternal := 0; indexExternal < 10; indexExternal++ {
		go func() {
			for index := 0; index < 10; index++ {
				channel <- index

			}
		}()
	}

	for index := 0; index < 100; index++ {
		fmt.Println(index, " Received ", <-channel)
	}
}
