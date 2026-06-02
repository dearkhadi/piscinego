package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	myVar := []int{}

	if n == 0 {
		z01.PrintRune('0')
	}

	for n > 0 {
		dig := n % 10
		myVar = append(myVar, dig)
		n = n / 10
	}

	for i := 0; i < len(myVar); i++ {
		for j := i + 1; j < len(myVar); j++ {
			if myVar[i] > myVar[j] {
				myVar[i], myVar[j] = myVar[j], myVar[i]
			}
		}
	}

	for _, dig := range myVar {
		z01.PrintRune(rune(dig) + '0')
	}
}
