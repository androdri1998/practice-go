package table

import (
	"strings"
	"testing"
)

const msgStandard = "%s (part: %s) - expected index(%d) <> found (%d)"

func TestIndex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Somethig about other thing", "other", 15},
		{"", "", 0},
		{"Hi, this is nice", "new", -1},
		{"someone", "o", 1},
	}

	for _, test := range tests {
		t.Logf("Data: %v", test)

		actualResult := strings.Index(test.text, test.part)

		if actualResult != test.expected {
			t.Errorf(msgStandard, test.text, test.part, test.expected, actualResult)
		}
	}
}
