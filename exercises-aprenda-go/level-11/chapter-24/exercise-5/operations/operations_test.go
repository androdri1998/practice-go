package operations_test

import (
	"testing"

	"github.com/androdri1998/practice-go/exercises-aprenda-go/level-11/chapter-24/exercise-5/operations"
)

func TestShouldSum(t *testing.T) {
	sum, _ := operations.Sum(1, 2)
	if sum != 3 {
		t.Fail()
	}
}

func TestShouldDivide(t *testing.T) {
	div, _ := operations.Division(4, 2)
	if div != 2 {
		t.Fail()
	}
}

func TestShouldNotDivide(t *testing.T) {
	_, err := operations.Division(4, 0)
	if err != nil {
		if err.Error() != "error: it's not possible divide 4 by 0 or less" {
			t.Fail()
		}
	}
}
