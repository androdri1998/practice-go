package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	values := `[{"First":"name 1","Last":"last name","Age":32,"Sayings":["Hello","World","In GO"]},{"First":"Name 2","Last":"lastname 2","Age":27,"Sayings":["Hello","World","in JavaScript"]},{"First":"Name 3","Last":"Lastname 3","Age":21,"Sayings":["Hello","World","In TypeScript"]}]`
	fmt.Println(values)

	var people []Person

	fmt.Println(people)
	json.Unmarshal([]byte(values), &people)
	fmt.Println(people)
}
