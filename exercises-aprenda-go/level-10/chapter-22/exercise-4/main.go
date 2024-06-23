package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	quit := make(chan bool)
	channel2 := genChannel(quit)

	wg.Add(1)
	receive(quit, channel2)

	wg.Wait()
	fmt.Println("End of execution")
}

func genChannel(quit chan<- bool) <-chan int {
	localChannel := make(chan int)

	go func() {
		for index := 0; index < 100; index++ {
			fmt.Println("Sending ", index)
			localChannel <- index
		}

		quit <- true

		close(localChannel)
	}()

	return localChannel
}

func receive(quit <-chan bool, channel <-chan int) {
	defer wg.Done()
	for {
		select {
		case value, ok := <-channel:
			if ok {
				fmt.Println("Receiving", value)
			}
		case value, ok := <-quit:
			if value && ok {
				fmt.Println("Quit execution")
				return
			}
		}
	}
}
