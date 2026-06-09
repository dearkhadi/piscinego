package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	var solve func(int)
	solve = func(col int) {
		if col == 8 {
			for i := 0; i < 8; i++ {
				z01.PrintRune(rune(board[i] + '0'))
			}
			z01.PrintRune('\n')
			return
		}
		for row := 1; row <= 8; row++ {
			safe := true
			for i := 0; i < col; i++ {
				rowDiff := board[i] - row
				if rowDiff < 0 {
					rowDiff = -rowDiff
				}
				colDiff := i - col
				if colDiff < 0 {
					colDiff = -colDiff
				}
				if board[i] == row || rowDiff == colDiff {
					safe = false
					break
				}
			}
			if safe {
				board[col] = row
				solve(col + 1)
			}
		}
	}
	solve(0)
}
