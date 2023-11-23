package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{10, 2, 4, 3, 100, 40, 30}
	strings := []string{"feliz", "alegre", "casa", "navio", "lua", "gato", "bota"}

	fmt.Println(ints)
	sort.Ints(ints)
	fmt.Println(ints)

	fmt.Println(strings)
	sort.Strings(strings)
	fmt.Println(strings)
}
