package main

import (
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n=1",
			n:    1,
			want: 1,
		},
		{
			name: "n=2",
			n:    2,
			want: 0,
		},
		{
			name: "n=4",
			n:    4,
			want: 2,
		},
		{
			name: "n=8",
			n:    8,
			want: 92,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solveNQueens(tt.n)
			if len(result) != tt.want {
				t.Errorf("got %d solutions, want %d", len(result), tt.want)
			}

			for _, board := range result {
				if len(board) != tt.n {
					t.Errorf("board has %d rows, want %d", len(board), tt.n)
				}
				for _, row := range board {
					if len(row) != tt.n {
						t.Errorf("row has length %d, want %d", len(row), tt.n)
					}
				}

				queenCount := 0
				for _, row := range board {
					for _, cell := range row {
						if cell == 'Q' {
							queenCount++
						}
					}
				}
				if queenCount != tt.n {
					t.Errorf("got %d queens, want %d", queenCount, tt.n)
				}

				validBoard := isValidNQueens(board, tt.n)
				if !validBoard {
					t.Errorf("board has queens attacking each other")
				}
			}
		})
	}
}

func isValidNQueens(board []string, n int) bool {
	positions := make([]int, n)
	for i, row := range board {
		for j, cell := range row {
			if cell == 'Q' {
				positions[i] = j
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if positions[i] == positions[j] {
				return false
			}
			if i-positions[i] == j-positions[j] {
				return false
			}
			if i+positions[i] == j+positions[j] {
				return false
			}
		}
	}
	return true
}
