package main

import "fmt"

type sportCar interface {
	turnOnTurbo()
}

type luxCar interface {
	autoParking()
}

type luxSportCar interface {
	sportCar
	luxCar
}

type carBrand struct{}

func (c carBrand) turnOnTurbo() {
	fmt.Println("Turbo on")
}

func (c carBrand) autoParking() {
	fmt.Println("Parking")
}

func main() {
	var car luxSportCar = carBrand{}

	car.turnOnTurbo()
	car.autoParking()

	fmt.Println(car)
}
