package main

import (
	"fmt"
	"time"
)

func rotinalCh(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1
	fmt.Println("Need to read channel...")
}

func main() {
	ch := make(chan int)

	go rotinalCh(ch)

	fmt.Println(<-ch)
	fmt.Println("Deadlock Example...")

	fmt.Println(<-ch)
	fmt.Println("End")
}
