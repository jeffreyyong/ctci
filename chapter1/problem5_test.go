package chapter1

import (
	"testing"
)

var oneEditAwayTests = []struct {
	input1   string
	input2   string
	expected bool
}{
	{"abcd", "abcd", true},
	{"abcd", "abcc", true},
	{"abcd", "accc", false},
	{"abcd", "abcde", true},
	{"abcd", "abcdef", false},
	{"abcde", "abcd", true},
	{"abcdef", "abcd", false},
	{" ", "", true},
	{"", " ", true},
	{"", "", true},
}

func TestAreOneEditAway(t *testing.T) {
	for _, c := range oneEditAwayTests {
		actual := AreOneEditAway(c.input1, c.input2)
		if actual != c.expected {
			t.Fatalf("Inputs %s, %s. Expected: %b, actual: %b\n", c.input1, c.input2, c.expected, actual)
		}
	}
}
