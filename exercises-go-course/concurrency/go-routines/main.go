package main

import (
	"fmt"
	"time"
)

func speak(person, text string, amount int) {
	for i := 0; i < amount; i++ {
		time.Sleep(time.Second)
		fmt.Println(person, text, i+1)
	}
}

func main() {
	// speak("joe", "Something 1", 3)
	// speak("doe", "Something 2", 1)

	// go speak("joe", "Something 1", 3)
	// go speak("doe", "Something 2", 1)

	go speak("joe", "Something 1", 10)
	speak("doe", "Something 2", 5)
	fmt.Println("End")
}
