package utils


func SortStr(s string) string{
	str := []byte(s)
	for i:= 0; i < len(str); i++{
		for j:= i+1; j < len(str); j++{
			if str[j] < str[i]{
				str[j], str[i] = str[i], str[j]
			}
		}
	}
	return string(str)
}
