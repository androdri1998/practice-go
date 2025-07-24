package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1
	channelValue := <-ch

	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(channelValue)
}
