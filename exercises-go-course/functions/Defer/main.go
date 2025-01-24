package main

import "fmt"

func getApprovedValue(number int) int {
	defer fmt.Println("End")
	if number > 5000 {
		fmt.Println("Too high")
		return 5000
	}

	fmt.Println("Too low")
	return number
}

func main() {
	getApprovedValue(5001)
	getApprovedValue(4099)
}
