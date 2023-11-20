package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}

type figure interface {
	area() float64
}

func info(fig figure) float64 {
	return fig.area()
}

func main() {
	newCircle := circle{2.4}
	newSquare := square{4.2}

	fmt.Println(info(newCircle))
	fmt.Println(info(newSquare))
}
