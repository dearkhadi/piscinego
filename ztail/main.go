package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		fmt.Fprintln(os.Stderr, "Usage: go run . -c <bytes> <file1> [file2 ...]")
		os.Exit(1)
	}

	var bytesToRead int64
	for _, ch := range os.Args[2] {
		if ch < '0' || ch > '9' {
			fmt.Fprintln(os.Stderr, "invalid number of bytes")
			os.Exit(1)
		}
		bytesToRead = bytesToRead*10 + int64(ch-'0')
	}

	files := os.Args[3:]
	multipleFiles := len(files) > 1
	hasError := false

	for i, filename := range files {
		if multipleFiles {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", filename)
		}

		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			hasError = true
			continue
		}

		stat, err := file.Stat()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			file.Close()
			hasError = true
			continue
		}

		fileSize := stat.Size()
		offset := fileSize - bytesToRead
		if offset < 0 {
			offset = 0
		}

		buffer := make([]byte, fileSize-offset)
		_, err = file.ReadAt(buffer, offset)
		file.Close()

		if err != nil && err.Error() != "EOF" {
			fmt.Fprintln(os.Stderr, err)
			hasError = true
			continue
		}

		fmt.Print(string(buffer))
	}

	if hasError {
		os.Exit(1)
	}
}
