package main

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(board, word, 0, i, j, make(map[[2]int]bool)) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, index int, row, col int, visited map[[2]int]bool) bool {
	if index == len(word) {
		return true
	}

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}

	pos := [2]int{row, col}
	if visited[pos] {
		return false
	}

	if board[row][col] != word[index] {
		return false
	}

	visited[pos] = true

	if dfs(board, word, index+1, row+1, col, visited) ||
		dfs(board, word, index+1, row-1, col, visited) ||
		dfs(board, word, index+1, row, col+1, visited) ||
		dfs(board, word, index+1, row, col-1, visited) {
		return true
	}

	delete(visited, pos)
	return false
}
