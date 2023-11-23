package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	user1 := user{
		First: "name 1",
		Age:   32,
	}

	user2 := user{
		First: "name 2",
		Age:   27,
	}

	user3 := user{
		First: "name 3",
		Age:   54,
	}

	users := []user{user1, user2, user3}

	fmt.Println(users)

	sb, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(sb))
}
