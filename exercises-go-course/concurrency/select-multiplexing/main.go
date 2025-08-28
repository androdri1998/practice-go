package main

import (
	"fmt"
	"time"
)

func speak(person string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("%s sepeaking: %d", person, i)
		}
	}()

	return ch
}

func join(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			select {
			case s := <-input1:
				ch <- s
			case s := <-input2:
				ch <- s
			}
		}
	}()

	return ch
}

func main() {
	ch := join(speak("Someone 1"), speak("Someone 2"))

	fmt.Println(<-ch, <-ch)
	fmt.Println(<-ch, <-ch)
	fmt.Println(<-ch, <-ch)
}
