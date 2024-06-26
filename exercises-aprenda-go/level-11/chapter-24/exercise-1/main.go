// remove underscore and check and handle error in a proper way
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

	valueInBytes, error := json.Marshal(p1)
	if error != nil {
		log.Println(error)
		return
	}

	fmt.Println(string(valueInBytes))
}
