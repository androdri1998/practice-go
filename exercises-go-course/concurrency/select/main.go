package main

import (
	"fmt"
	"time"

	p "github.com/androdri1998/practice-go/exercises-go-course/concurrency/multiplexing/places"
)

func faster(runner1, runner2, runner3 string) string {
	ch1 := p.Places(runner1)
	ch2 := p.Places(runner2)
	ch3 := p.Places(runner3)

	select {
	case r1 := <-ch1:
		return r1
	case r2 := <-ch2:
		return r2
	case r3 := <-ch3:
		return r3
	case <-time.After(4000 * time.Millisecond):
		return "All lose the run"
	}
}

func main() {
	champion := faster("Runner 1", "Runner 2", "Runner 3")
	fmt.Println(champion)
}
