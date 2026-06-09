package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	size := len("........")
	one := len("a")
	zero := len("")

	var solve func(int)
	solve = func(col int) {
		if col == size {
			for i := zero; i < size; i++ {
				z01.PrintRune(rune(board[i] + '0'))
			}
			z01.PrintRune('\n')
			return
		}
		for row := one; row <= size; row++ {
			safe := true
			for i := zero; i < col; i++ {
				rowDiff := board[i] - row
				if rowDiff < zero {
					rowDiff = -rowDiff
				}
				colDiff := i - col
				if colDiff < zero {
					colDiff = -colDiff
				}
				if board[i] == row || rowDiff == colDiff {
					safe = false
					break
				}
			}
			if safe {
				board[col] = row
				solve(col + one)
			}
		}
	}
	solve(zero)
}
