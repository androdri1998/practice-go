package main

import "fmt"

func main() {
	fmt.Println(multiplyNumberBy2(10))
	fmt.Println(getLetters(10, "a"))
}

func multiplyNumberBy2(number int) int {
	return number * 2
}

func getLetters(number int, letter string) (int, string) {
	phrase := ""
	for index := 0; index < number; index++ {
		phrase = fmt.Sprintf("%v%v", phrase, letter)
	}

	return number, phrase
}
