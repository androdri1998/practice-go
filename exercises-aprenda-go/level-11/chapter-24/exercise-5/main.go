// create a function
// create an unit test for this function in go lang
package main

import (
	"fmt"
	"log"

	"github.com/androdri1998/practice-go/exercises-aprenda-go/level-11/chapter-24/exercise-5/operations"
)

func main() {
	valueSum, _ := operations.Sum(3, 4)
	fmt.Println(valueSum)

	valueDiv1, errFirstDiv := operations.Division(4, 2)
	if errFirstDiv != nil {
		log.Println(errFirstDiv)
		return
	}
	fmt.Println(valueDiv1)

	valueDiv2, errSecondDiv := operations.Division(4, 0)
	if errSecondDiv != nil {
		log.Println(errSecondDiv)
		return
	}
	fmt.Println(valueDiv2)
}
