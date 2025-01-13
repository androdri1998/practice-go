package main

import "fmt"

func printGrade(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved with grade:", grade)
		return
	}

	fmt.Println("Reproved with grade:", grade)
	return
}

func main() {
	printGrade(7.2)
	printGrade(6.5)
}
