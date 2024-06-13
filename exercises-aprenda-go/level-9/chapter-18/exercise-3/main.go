// create a counter
// have a variable with value
// create a go routine, that need:
// - read value from counter
// - save value
// - to do a yield with runtime.Gosched
// - increment value
// - copy value to first variable
// use Waitgroups
// need to run go run -race main.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int64 = 0
	limit := 100

	wg.Add(limit)

	for index := 0; index < limit; index++ {
		go func() {
			var localCounter int64 = counter
			runtime.Gosched()

			localCounter++

			counter = localCounter

			wg.Done()
		}()
	}

	fmt.Println("Final value", counter)

	wg.Wait()
}
