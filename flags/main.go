package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		for _, r := range "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n" {
			z01.PrintRune(r)
		}
		return
	}

	insertStr := ""
	hasInsert := false
	hasOrder := false
	baseStr := ""
	hasBase := false

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			for _, r := range "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n" {
				z01.PrintRune(r)
			}
			return
		} else if arg == "-o" || arg == "--order" {
			hasOrder = true
		} else if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
			hasInsert = true
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
			hasInsert = true
		} else if !hasBase {
			baseStr = arg
			hasBase = true
		}
	}

	result := baseStr
	if hasInsert {
		result += insertStr
	}

	if hasOrder {
		runes := []rune(result)
		for i := 0; i < len(runes); i++ {
			for j := i + 1; j < len(runes); j++ {
				if runes[i] > runes[j] {
					runes[i], runes[j] = runes[j], runes[i]
				}
			}
		}
		result = string(runes)
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
