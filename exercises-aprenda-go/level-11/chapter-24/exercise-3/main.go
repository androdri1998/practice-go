package main

import (
	"fmt"
	"log"
)

type specialError struct {
	desc string
	err  error
}

func (err specialError) Error() string {
	return fmt.Sprintf("Message: %v", err.desc)
}

func main() {
	value, err := fmt.Println("Some value")
	if err != nil {
		newError := specialError{desc: "this is the error", err: err}
		logError(newError)
	}

	fmt.Println(value)
}

func logError(err error) {
	log.Println(err)
}
