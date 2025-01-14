package main

import "fmt"

func convertGradeToConcept(grade float64) string {
	gradeInt := int(grade)

	switch gradeInt {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
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
