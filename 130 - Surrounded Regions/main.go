package main

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (i == 0 || i == rows-1 || j == 0 || j == cols-1) && board[i][j] == 'O' {
				dfs(board, i, j, rows, cols)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j, rows, cols int) {
	if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != 'O' {
		return
	}

	board[i][j] = '#'
	dfs(board, i+1, j, rows, cols)
	dfs(board, i-1, j, rows, cols)
	dfs(board, i, j+1, rows, cols)
	dfs(board, i, j-1, rows, cols)
}
