package main

import "fmt"

type carBrandSport interface {
	turnOnTurbo()
}

type carBrand struct {
	model        string
	isTurboOn    bool
	currentSpeed int
}

func (c *carBrand) turnOnTurbo() {
	c.isTurboOn = true
}

func main() {
	var car1 carBrandSport = &carBrand{"One car", false, 0}
	car1.turnOnTurbo()
	fmt.Println(car1)

	car2 := carBrand{"One car", false, 0}
	car2.turnOnTurbo()
	fmt.Println(car2)
}
