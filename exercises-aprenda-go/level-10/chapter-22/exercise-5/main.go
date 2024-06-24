package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 20
	}()

	value, ok := <-channel
	fmt.Println(value, ok)

	close(channel)

	value, ok = <-channel
	fmt.Println(value, ok)
}
