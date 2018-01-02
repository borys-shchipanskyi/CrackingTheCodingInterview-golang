package chapter1

import "testing"

func TestProblem1(t *testing.T) {
	tests := []struct {
		input string
		out   bool
	}{
		{"abc", true},
		{"", true},
		{"aa", false},
		{"abca", false},
	}

	for _, test := range tests {

		res := isUnique(test.input)
		if res != test.out {
			t.Errorf("Test input: %#v, Expect: %v, but get: %v\n", test.input, test.out, res)
		}
	}
}
