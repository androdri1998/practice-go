package places

import (
	"fmt"
	"math/rand"
	"time"
)

func Places(runners ...string) <-chan string {
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
