package main

import (
	"fmt"
	"time"
)

func main() {
	currentYear, _, _ := time.Now().Date()

	for year := 1998; year <= currentYear; year++ {
		fmt.Println(year)
	}
}
