package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0
	for _, value := range numbers {
		total += value
	}

	return total / float64(len(numbers))
}

func main() {
	fmt.Println(average(1, 2, 3, 4, 5))
}
