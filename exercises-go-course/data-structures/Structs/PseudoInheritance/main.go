package main

import "fmt"

type car struct {
	name         string
	currentSpeed int
}

type carBrand struct {
	car
	turnOnTurbo bool
}

func main() {
	c := carBrand{}

	c.name = "Some car"
	c.currentSpeed = 0
	c.turnOnTurbo = true

	fmt.Println(c, c.name, c.currentSpeed, c.turnOnTurbo)
}
