package main

import (
	"fmt"

	"github.com/androdri1998/practice-go/exercises-aprenda-go/level-12/chapter-26/exercise-1/dog"
)

func main() {
	dogYears := dog.ConvertYear(12)
	fmt.Println(12, " years in dog years it's equivalent to ", dogYears)
}
