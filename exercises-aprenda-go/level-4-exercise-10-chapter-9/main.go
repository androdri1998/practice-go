package main

import "fmt"

func main() {
	hashTable := map[string][]string{
		"name_lastname1": {"hobbie 1", "hobbie 2", "hobbie 3"},
		"name_lastname2": {"hobbie 4", "hobbie 5", "hobbie 6"},
		"name_lastname3": {"hobbie 7", "hobbie 8", "hobbie 9"},
	}

	hashTable["newname_lastname"] = []string{"hobbie 10", "hobbie 11", "hobbie 12"}

	for person, hobbies := range hashTable {
		println(person)
		for _, hobby := range hobbies {
			fmt.Println("\t", hobby)
		}
	}

	delete(hashTable, "newname_lastname")
	println("-------------")

	for person, hobbies := range hashTable {
		println(person)
		for _, hobby := range hobbies {
			fmt.Println("\t", hobby)
		}
	}
}
