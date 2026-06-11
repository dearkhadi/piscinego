package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	n := 0

	for _, r := range s {
		n = n*10 + int(r-'0')
	}

	return n
}

func main() {
	if len(os.Args) < 4 {
		return
	}

	count := Atoi(os.Args[2])
	files := os.Args[3:]

	hasError := false

	for i, file := range files {
		data, err := os.ReadFile(file)

		if err != nil {
			fmt.Printf("%v\n", err)
			hasError = true
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", file)
		}

		start := 0
		if count < len(data) {
			start = len(data) - count
		}

		fmt.Printf("%s", data[start:])
	}

	if hasError {
		os.Exit(1)
	}
}
