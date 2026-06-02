package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]

	if len(arguments) > 1 && arguments[0] == '.' && arguments[1] == '/' {
		arguments = arguments[2:]
	}

	for _, symbol := range arguments {
		z01.PrintRune(symbol)
	}
	z01.PrintRune('\n')
}
