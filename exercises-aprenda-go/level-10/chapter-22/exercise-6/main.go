package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for index := 0; index < 100; index++ {
			channel <- index
		}
	}()

	for value := range channel {
		fmt.Println("Received ", value+1)
	}
}
