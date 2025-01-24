package main

import "fmt"

func printApproved(approvedList ...string) {
	for _, value := range approvedList {
		fmt.Println(value)
	}
}

func main() {
	approvedList := []string{"Joe", "Doe"}
	printApproved(approvedList...)
}
