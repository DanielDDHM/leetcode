package main

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			char := board[i][j]
			if char == '.' {
				continue
			}

			boxIdx := (i/3)*3 + j/3

			if rows[i][char] || cols[j][char] || boxes[boxIdx][char] {
				return false
			}

			rows[i][char] = true
			cols[j][char] = true
			boxes[boxIdx][char] = true
		}
	}

	return true
}
