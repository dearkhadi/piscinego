package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	myRunes := []rune{}

	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if !(i == 99 && j == 98) {
				myRunes = append(myRunes, ',', ' ')
			}
			myRunes = append(myRunes, '0'+rune(i/10), '0'+rune(i%10), ' ', '0'+rune(j/10), '0'+rune(j%10))
		}
	}
	for _, rune := range myRunes {
		z01.PrintRune(rune)
	}
}
