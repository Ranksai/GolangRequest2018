package isEven

import (
	"testing"
)

func TestIsEven(t *testing.T) {
	t.Log("TestIsEven")

	cases := []struct {
		name     string
		input    int
		expected bool
	}{
		{name: "+odd", input: 5, expected: false},
		{name: "+even", input: 6, expected: true},
		{name: "-odd", input: -5, expected: false},
		{name: "-even", input: -6, expected: true},
		{name: "zero", input: 0, expected: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := IsEven(c.input); c.expected != actual {
				t.Errorf("want IsEven(%d) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}
