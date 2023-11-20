package main

import "fmt"

func main() {
	slice := []string{"one", "two", "three", "four", "five"}

	for index := 0; index < len(slice); index++ {
		fmt.Println(slice[index])
	}
}
