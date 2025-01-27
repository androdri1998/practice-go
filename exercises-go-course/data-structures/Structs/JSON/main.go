package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := product{1, "Laptop", 1800.90, []string{"Sales", "Gadget"}}
	p1JSON, _ := json.Marshal(p1)

	fmt.Println(string(p1JSON))

	var p2 product
	jsonString := `{"id":2,"name":"Pen","price":8.90,"tags":["Paper","Import"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2, p2.Tags, p2.Tags[0])
}
