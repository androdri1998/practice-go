package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ageSort []user
type lastSort []user

func (users ageSort) Len() int {
	return len(users)
}

func (users ageSort) Less(a, b int) bool {
	return users[a].Age < users[b].Age
}

func (users ageSort) Swap(a, b int) {
	users[a], users[b] = users[b], users[a]
}

func (users lastSort) Len() int {
	return len(users)
}

func (users lastSort) Less(a, b int) bool {
	return users[a].Last < users[b].Last
}

func (users lastSort) Swap(a, b int) {
	users[a], users[b] = users[b], users[a]
}

func main() {
	user1 := user{
		First: "name 1",
		Age:   32,
		Last:  "lastname 1",
		Sayings: []string{
			"Hello",
			"World",
			"In GO",
		},
	}

	user2 := user{
		First: "name 2",
		Age:   27,
		Last:  "lastname 2",
		Sayings: []string{
			"Hello",
			"World",
			"In JavaScript",
		},
	}

	user3 := user{
		First: "name 3",
		Age:   54,
		Last:  "lastname 3",
		Sayings: []string{
			"Hello",
			"World",
			"In TypeScript",
		},
	}

	users := ageSort{user3, user1, user2}

	fmt.Println(users)
	sort.Sort(ageSort(users))
	fmt.Println(users)
	sort.Sort(lastSort(users))
	fmt.Println(users)

	for _, value := range users {
		sort.Strings(value.Sayings)

		fmt.Println("")
		fmt.Println(value.First)
		fmt.Println(value.Sayings)
	}
}
