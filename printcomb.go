package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := '1'; b <= '9'; b++ {
			if b > a {
				for c := '2'; c <= '9'; c++ {
					if c > b {
						if a != '7' || b != '8' || c != '9' {
							z01.PrintRune(a)
							z01.PrintRune(b)
							z01.PrintRune(c)
							z01.PrintRune(',')
							z01.PrintRune(' ')
						} else {
							z01.PrintRune(a)
							z01.PrintRune(b)
							z01.PrintRune(c)
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
