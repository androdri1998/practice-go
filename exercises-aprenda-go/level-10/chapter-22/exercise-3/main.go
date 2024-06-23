package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	channel := genChannel()

	wg.Add(1)
	go receive(channel)

	wg.Wait()
	fmt.Println("End of execution")
}

func genChannel() chan int {
	channel := make(chan int, 100)

	go func() {
		defer close(channel)
		for index := 0; index < 100; index++ {
			fmt.Println("Sending ", index)
			channel <- index
		}
	}()

	return channel
}

func receive(channel <-chan int) {
	defer wg.Done()
	for value := range channel {
		fmt.Println("Receiving ", value)
	}
}
