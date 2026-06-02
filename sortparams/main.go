package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	for i := 1; i < len(arguments); i++ {
		for j := i + 1; j < len(arguments); j++ {
			if arguments[i] > arguments[j] {
				arguments[i], arguments[j] = arguments[j], arguments[i]
			}
		}
	}

	for i := 1; i < len(arguments); i++ {
		for _, symbol := range arguments[i] {
			z01.PrintRune(symbol)
		}
		z01.PrintRune('\n')
	}
}
