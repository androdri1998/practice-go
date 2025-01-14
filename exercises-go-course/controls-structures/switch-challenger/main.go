package main

import "fmt"

func convertGradeToConcept(grade float64) string {
	gradeInt := int(grade)

	switch {
	case gradeInt >= 9 && gradeInt <= 10:
		return "A"
	case gradeInt >= 7 && gradeInt < 9:
		return "B"
	case gradeInt >= 5 && gradeInt < 7:
		return "C"
	case gradeInt >= 3 && gradeInt < 5:
		return "D"
	case gradeInt >= 0 && gradeInt < 3:
		return "E"
	default:
		return "Not supported"
	}
}

func main() {
	fmt.Println(convertGradeToConcept(9.8))
	fmt.Println(convertGradeToConcept(8.3))
	fmt.Println(convertGradeToConcept(6.4))
	fmt.Println(convertGradeToConcept(4.3))
	fmt.Println(convertGradeToConcept(2.3))
}
