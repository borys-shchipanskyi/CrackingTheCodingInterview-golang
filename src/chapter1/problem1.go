package chapter1

// Problem 1.1 Is Unique: Implementing an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structure?

func isUnique(s string) bool {
	for i, _ := range s{
		for j:= i+1; j<len(s); j++{
			if s[i] == s[j]{
				return false
			}
		}
	}
	return true
}
