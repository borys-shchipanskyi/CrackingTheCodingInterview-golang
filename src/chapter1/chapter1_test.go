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

func TestProblem2(t *testing.T) {
	tests := []struct {
		s1 string
		s2 string
		out   bool
	}{
		{"abc", "abc",  false},
		{"abc", "acb",  true},
		{"abc", "Abc",  false},
		{"", "c",  false},
	}

	for _, test := range tests {

		res := isPermutation(test.s1, test.s2)
		if res != test.out {
			t.Errorf("Test s1: %#v, s2: %#v, Expect: %v, but get: %v\n", test.s1, test.s2, test.out, res)
		}
	}
}