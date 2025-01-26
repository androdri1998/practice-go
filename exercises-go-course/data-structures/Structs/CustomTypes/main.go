package main

import "fmt"

type grade float64

func (g grade) between(start, end float64) bool {
	return float64(g) >= start && float64(g) <= end
}

func gradeToConcept(g grade) string {
	if g.between(9.0, 10.0) {
		return "A"
	}

	if g.between(7.0, 8.99) {
		return "B"
	}

	if g.between(5.0, 7.99) {
		return "C"
	}

	if g.between(3.0, 4.99) {
		return "D"
	}

	return "E"
}

func main() {
	fmt.Println(gradeToConcept(9.8))
	fmt.Println(gradeToConcept(7.5))
	fmt.Println(gradeToConcept(5.6))
	fmt.Println(gradeToConcept(4.2))
	fmt.Println(gradeToConcept(2.1))
}
