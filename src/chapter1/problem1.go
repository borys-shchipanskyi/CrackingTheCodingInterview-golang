package chapter1

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
