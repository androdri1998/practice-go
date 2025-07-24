package main

import (
	"fmt"
	"time"
)

func multiplyBy(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	ch := make(chan int)
	go multiplyBy(2, ch)
	fmt.Println("Storing...")

	a, b := <-ch, <-ch
	fmt.Println("Reading...")
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(<-ch)
}
