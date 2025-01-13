package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return int(r.Intn(10))
}

func main() {
	if i := randomNumber(); i > 5 {
		fmt.Println("Won:", i)
	} else {
		fmt.Println("lost:", i)
	}
}
