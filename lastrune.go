package piscine

func LastRune(s string) rune {
	myVar := []rune(s)

	return myVar[len(myVar)-1]
}
