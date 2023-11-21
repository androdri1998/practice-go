package main

import "fmt"

func main() {
	person := struct {
		things  map[string]int
		hobbies []string
	}{
		things: map[string]int{
			"book":  2,
			"games": 1,
		},
		hobbies: []string{"play soccer", "read book", "play video games"},
	}

	fmt.Println(person)
}
