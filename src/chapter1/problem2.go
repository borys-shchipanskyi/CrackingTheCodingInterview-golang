package chapter1

import . "utils"

// Problem 1.2 Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.

// I expect that "permutation" meant than strings are case sensitive and are not equal.
func isPermutation(s1, s2 string) bool{
	if len(s1) != len(s2) || s1 == s2{
		return false
	}
	s1 = SortStr(s1)
	s2 = SortStr(s2)
	for i := 0; i < len(s1); i++{
		if s1[i] != s2[i]{
			return false
		}
	}
	return true
}