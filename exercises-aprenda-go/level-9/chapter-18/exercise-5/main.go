package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter int64 = 0
	limit := 100

	wg.Add(limit)

	for index := 0; index < limit; index++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			localCounter := atomic.LoadInt64(&counter)
			fmt.Println(localCounter)

			runtime.Gosched()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final value", counter)
}
