package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	upper := false
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		if len(arg) == 0 {
			z01.PrintRune(' ')
			continue
		}

		n := 0
		isValid := true
		for i := 0; i < len(arg); i++ {
			if arg[i] < '0' || arg[i] > '9' {
				isValid = false
				break
			}
			n = n*10 + int(arg[i]-'0')
			if n > 26 {
				isValid = false
				break
			}
		}

		if !isValid || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		if upper {
			z01.PrintRune(rune('A' + n - 1))
		} else {
			z01.PrintRune(rune('a' + n - 1))
		}
	}
	z01.PrintRune('\n')
}
