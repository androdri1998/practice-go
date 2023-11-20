package main

import "fmt"

func main() {
	favoriteSport := "soccer"

	switch favoriteSport {
	case "soccer":
		fmt.Printf("your favorite sport it's soccer")
	case "Football":
		fmt.Printf("your favorite sport it's Football")
	default:
		fmt.Printf("you have another favorite sport")
	}
}
