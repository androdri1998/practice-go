package main

import "fmt"

func evaluateGrade(grade float64) string {
	if grade >= 6 {
		return "Approved"
	}

	return "repproved"
}

func main() {
	grade1 := evaluateGrade(6.1)
	grade2 := evaluateGrade(6.0)
	grade3 := evaluateGrade(5.9)

	fmt.Println(grade1, grade2, grade3)
}
