package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	isFirst := true
	combination := make([]byte, n)
	generateComb(0, 0, n, combination, &isFirst)
	z01.PrintRune('\n')
}

func generateComb(currentDigit int, index int, n int, comb []byte, isFirst *bool) {
	if index == n {
		if !*isFirst {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		*isFirst = false

		for i := 0; i < n; i++ {
			z01.PrintRune(rune(comb[i]))
		}
		return
	}

	for i := currentDigit; i <= 9; i++ {
		comb[index] = byte('0' + i)
		generateComb(i+1, index+1, n, comb, isFirst)
	}
}
