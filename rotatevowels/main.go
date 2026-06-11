package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	var vowels []rune
	for _, arg := range args {
		for _, r := range arg {
			if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
				r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
				vowels = append(vowels, r)
			}
		}
	}

	vIdx := len(vowels) - 1
	for i, arg := range args {
		for _, r := range arg {
			if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
				r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
				z01.PrintRune(vowels[vIdx])
				vIdx--
			} else {
				z01.PrintRune(r)
			}
		}
		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
