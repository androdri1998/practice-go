package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	limit := 100

	wg.Add(limit)
	for index := 0; index < limit; index++ {
		go func() {
			mutex.Lock()
			localCounter := counter
			runtime.Gosched()

			localCounter++

			counter = localCounter
			mutex.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Final value", counter)
}
