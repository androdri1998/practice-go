package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitFunctions sync.WaitGroup

func main() {
	fmt.Printf("%d go routines running\n", runtime.NumGoroutine())

	waitFunctions.Add(2)

	fmt.Printf("functions called\n")
	go firstCounter()
	go secondCounter()

	fmt.Printf("%d go routines running\n", runtime.NumGoroutine())

	waitFunctions.Wait()
	fmt.Printf("done\n")
}

func firstCounter() {
	for index := 1; index <= 1000000; index++ {
		fmt.Printf("First counter %d \n", index)
	}

	waitFunctions.Done()
}

func secondCounter() {
	for index := 1; index <= 1000000; index++ {
		fmt.Printf("Second counter %d \n", index)
	}

	waitFunctions.Done()
}
