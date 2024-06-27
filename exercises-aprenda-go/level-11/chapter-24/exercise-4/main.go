package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat string
	lon string
	err error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.lon, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}

}

func sqrt(number float64) (float64, error) {
	if number < 0 {
		err := sqrtError{lat: "some lat", lon: "some lon", err: errors.New("Error: it's not possible to do calc")}
		return number, err
	}

	return 20, nil
}
