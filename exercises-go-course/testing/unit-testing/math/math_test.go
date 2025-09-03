package math

import "testing"

const standardError = "Expected value %v, but actual value is %v"

func TestAverage(t *testing.T) {
	t.Parallel()
	expectedValue := 2.50
	actualValue := Average(1.5, 2.5, 1.4, 4.6)

	if actualValue != expectedValue {
		t.Errorf(standardError, expectedValue, actualValue)
	}
}
