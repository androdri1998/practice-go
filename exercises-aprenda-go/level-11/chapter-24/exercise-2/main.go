package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName string
	LastName  string
	Sayings   []string
}

func main() {
	p1 := person{
		FirstName: "Some",
		LastName:  "one",
		Sayings:   []string{"just", "saying", "something"},
	}

	valueInBytes, err := toJSON(p1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(valueInBytes))
}

func toJSON(p person) ([]byte, error) {

	valueInBytes, err := json.Marshal(p)

	if err != nil {
		return []byte{}, fmt.Errorf("something wrong happen")
	}

	return valueInBytes, nil
}
