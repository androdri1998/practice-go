package main

import "fmt"

func main() {
	matrix := [][]string{}

	matrix = append(matrix, []string{"name 1", "lastname 1", "soccer"})
	matrix = append(matrix, []string{"name 2", "lastname 2", "voleyball"})
	matrix = append(matrix, []string{"name 3", "lastname 3", "basketball"})

	for _, value := range matrix {
		fmt.Println(value[0], value[1], value[2])
	}
}
