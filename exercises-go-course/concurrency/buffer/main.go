package main

import "fmt"

func routineCh(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Running routine...")
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go routineCh(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
