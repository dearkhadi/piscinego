package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		_, _ = io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {
			printStr("ERROR: ")
			printStr(err.Error())
			printStr("\n")
			os.Exit(1)
		}

		_, _ = io.Copy(os.Stdout, file)
		file.Close()
	}
}
