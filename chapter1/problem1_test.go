package chapter1

import (
	"testing"
)

var testCases = []struct {
	input    string
	expected bool
}{
	{"abcd", true},
	{"abcc", false},
	{" ", true},
	{"", true},
}

func TestIsUnique(t *testing.T) {
	for _, c := range testCases {
		actual := IsUnique(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %b, actual: %b\n", c.input, c.expected, actual)
		}
	}
}
