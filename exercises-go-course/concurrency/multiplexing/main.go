package main

import (
	"fmt"

	p "github.com/androdri1998/practice-go/exercises-go-course/concurrency/multiplexing/places"
)

func forwardMessage(source <-chan string, receiver chan string) {
	for {
		receiver <- <-source
	}
}

func join(source1, source2 <-chan string) <-chan string {
	ch := make(chan string)
	go forwardMessage(source1, ch)
	go forwardMessage(source2, ch)

	return ch
}

func main() {
	ch := join(
		p.Places("runner 1", "runner 2"),
		p.Places("runner 3", "runner 4"),
	)

	fmt.Println(<-ch, "|", <-ch)
	fmt.Println(<-ch, "|", <-ch)
}
