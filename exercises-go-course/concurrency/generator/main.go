package main

import (
	"fmt"
	"math/rand"
	"time"
)

func places(runners ...string) <-chan string {
	ch := make(chan string)
	for _, runner := range runners {
		go func(runner string) {
			randomSecond := rand.Intn(10) + 1
			time.Sleep(time.Duration(randomSecond) * time.Second)

			ch <- fmt.Sprintf("%s - %v seconds", runner, randomSecond)
		}(runner)
	}

	return ch
}

func main() {
	p1 := places("runner 1", "runner 2")
	p2 := places("runner 3", "runner 4")

	fmt.Println("First places: ", <-p1, "|", <-p2)
	fmt.Println("Runners-up: ", <-p1, "|", <-p2)
}
