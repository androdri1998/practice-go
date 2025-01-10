package helpers

import "fmt"

func Sum(a int64, b int64) (int64, error) {
	result := a + b

	return result, nil
}

func PrintInt(value int64) {
	fmt.Println(value)
}

func PrintError(value error) {
	fmt.Println(value)
}
