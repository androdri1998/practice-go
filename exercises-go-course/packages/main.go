package main

import (
	"fmt"

	s "github.com/androdri1998/practice-go/exercises-go-course/packages/straight"
	a "github.com/androdri1998/testgolibarea"
)

func main() {
	p1 :=
		s.Point{X: 2.0, Y: 2.0}
	p2 :=
		s.Point{X: 2.0, Y: 4.0}

	fmt.Println(s.Peccaries(p1, p2))
	fmt.Println(s.Distance(p1, p2))

	fmt.Println(a.Circle(10))
	fmt.Println(a.Rect(2, 4))
}
