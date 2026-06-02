package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	argsCount := len(os.Args) - 1

	if argsCount < 1 {
		fmt.Println("File name missing")
		return
	}
	if argsCount > 1 {
		fmt.Println("Too many arguments")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	_, _ = io.Copy(os.Stdout, file)
}
