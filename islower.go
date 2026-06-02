package piscine

func IsLower(s string) bool {
	myVar := []rune(s)
	answer := false

	for _, symbol := range myVar {
		if symbol >= 'a' && symbol <= 'z' {
			answer = true
		} else {
			answer = false
			break
		}
	}
	return answer
}
