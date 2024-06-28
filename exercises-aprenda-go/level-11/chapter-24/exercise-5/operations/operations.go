package operations

import (
	"fmt"
)

func Sum(a float64, b float64) (float64, error) {
	sum := a + b

	return sum, nil
}

func Division(a float64, b float64) (float64, error) {
	if b <= 0 {
		return 0, fmt.Errorf("error: it's not possible divide %v by 0 or less", a)
	}

	division := a / b

	return division, nil
}
