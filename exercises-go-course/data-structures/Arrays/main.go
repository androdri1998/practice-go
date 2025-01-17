package main

import "fmt"

func main() {
	// array only contains the same type of data and do not change the length
	var grades [3]float64
	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 7.4, 4.2, 9.3
	fmt.Println(grades)

	var sum float64 = 0
	for _, grade := range grades {
		sum += grade
	}

	average := sum / float64(len(grades))
	fmt.Printf("Average: %.2f\n", average)
}
