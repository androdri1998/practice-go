package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	user1 := user{
		First: "name 1",
		Age:   32,
		Sayings: []string{
			"Hello",
			"World",
			"In GO",
		},
	}

	user2 := user{
		First: "name 2",
		Age:   27,
		Sayings: []string{
			"Hello",
			"World",
			"In JavaScript",
		},
	}

	user3 := user{
		First: "name 3",
		Age:   54,
		Sayings: []string{
			"Hello",
			"World",
			"In TypeScript",
		},
	}

	users := []user{user1, user2, user3}

	fmt.Println(users)
	json.NewEncoder(os.Stdout).Encode(users)
}
