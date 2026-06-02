package piscine

func AlphaCount(s string) int {
	var counter int
	counter = 0
	myVar := []rune(s)

	for _, symbol := range myVar {
		if (symbol >= 'A' && symbol <= 'Z') || (symbol >= 'a' && symbol <= 'z') {
			counter++
		}
	}
	return counter
}
