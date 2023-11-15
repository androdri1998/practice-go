package main

import (
	"fmt"
	"time"
)

func main() {
	currentYear, _, _ := time.Now().Date()

	year := 1998
	for {
		fmt.Println(year)
		year++

		if year > currentYear {
			break
		}
	}
}
