package main

import (
	"fmt"
	"time"
)

func isDivisible(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func divisibleNumbers(n int, ch chan int) {
	start := 2
	for i := 0; i < n; i++ {
		for number := start; ; number++ {
			if isDivisible(number) {
				ch <- number
				start = number + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}

	close(ch)
}

func main() {
	ch := make(chan int, 30)

	go divisibleNumbers(cap(ch), ch)

	for number := range ch {
		fmt.Printf("%d ", number)
	}

	fmt.Println("End")
}
