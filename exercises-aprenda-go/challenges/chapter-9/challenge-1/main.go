package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	total := sumOddNumbers(sum, numbers)

	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func sumOddNumbers(sum func(...int) int, values []int) int {
	oddNumbers := []int{}
	for _, number := range values {
		if number%2 != 0 {
			oddNumbers = append(oddNumbers, number)
		}
	}

	total := sum(oddNumbers...)
	return total
}
