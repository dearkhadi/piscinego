package piscine

func IsAlpha(s string) bool {
	myVar := []rune(s)
	answer := false

	for _, symbol := range myVar {
		if (symbol >= 'a' && symbol <= 'z') ||
			(symbol >= 'A' && symbol <= 'Z') ||
			(symbol >= '0' && symbol <= '9') {
			answer = true
		} else {
			answer = false
			break
		}
	}
	return answer
}
