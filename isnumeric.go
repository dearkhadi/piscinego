package piscine

func IsNumeric(s string) bool {
	myVar := []rune(s)
	answer := false

	for _, symbol := range myVar {
		if symbol >= '0' && symbol <= '9' {
			answer = true
		} else {
			answer = false
			break
		}
	}
	return answer
}
