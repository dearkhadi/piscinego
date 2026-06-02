package piscine

func ToUpper(s string) string {
	myVar := []rune(s)

	for index, symbol := range myVar {
		if symbol >= 'a' && symbol <= 'z' {
			myVar[index] = symbol - 32
		}
	}
	s = string(myVar)
	return s
}
