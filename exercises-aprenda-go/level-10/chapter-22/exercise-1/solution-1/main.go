package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 20
	}()

	fmt.Println(<-channel)
}
