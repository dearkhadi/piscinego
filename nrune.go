package piscine

func NRune(s string, n int) rune {
	myVar := []rune(s)

	if n > len(myVar) || n <= 0 {
		return 0
	} else {
		return myVar[n-1]
	}
}
