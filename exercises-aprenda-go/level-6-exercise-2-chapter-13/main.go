package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(sumValues(numbers...))
	fmt.Println(sumValuesSlice(numbers))
}

func sumValues(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}

func sumValuesSlice(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}
