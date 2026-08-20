package main

func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			solution := make([]string, n)
			for i := range board {
				solution[i] = string(board[i])
			}
			result = append(result, solution)
			return
		}

		for col := 0; col < n; col++ {
			d1 := row - col + n - 1
			d2 := row + col

			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}

			board[row][col] = 'Q'
			cols[col] = true
			diag1[d1] = true
			diag2[d2] = true

			backtrack(row + 1)

			board[row][col] = '.'
			cols[col] = false
			diag1[d1] = false
			diag2[d2] = false
		}
	}

	backtrack(0)
	return result
}
