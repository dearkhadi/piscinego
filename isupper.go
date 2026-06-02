package piscine

func IsUpper(s string) bool {
	myVar := []rune(s)
	answer := false
	for _, symbol := range myVar {
		if symbol >= 'A' && symbol <= 'Z' {
			answer = true
		} else {
			answer = false
			break
		}
	}
	return answer
}
