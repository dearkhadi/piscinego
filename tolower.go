package piscine

func ToLower(s string) string {
	myVar := []rune(s)

	for index, symbol := range myVar {
		if symbol >= 'A' && symbol <= 'Z' {
			myVar[index] = symbol + 32
		}
	}
	s = string(myVar)
	return s
}
