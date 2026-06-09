package piscine

import "fmt"

func EightQueens() {
	var board [8]int
	solve(0, &board)
}

func solve(col int, board *[8]int) {
	if col == 8 {
		for i := 0; i < 8; i++ {
			fmt.Print(board[i])
		}
		fmt.Println()
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(col, row, board) {
			board[col] = row
			solve(col+1, board)
		}
	}
}

func isSafe(col, row int, board *[8]int) bool {
	for i := 0; i < col; i++ {
		prevRow := board[i]
		if prevRow == row {
			return false
		}
		rowDiff := prevRow - row
		if rowDiff < 0 {
			rowDiff = -rowDiff
		}
		colDiff := i - col
		if colDiff < 0 {
			colDiff = -colDiff
		}
		if rowDiff == colDiff {
			return false
		}
	}
	return true
}
