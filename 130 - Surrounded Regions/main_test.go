package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		expected [][]byte
	}{
		{
			name: "example 1",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "all X",
			board: [][]byte{
				{'X', 'X'},
				{'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X'},
				{'X', 'X'},
			},
		},
		{
			name: "O at edges not captured",
			board: [][]byte{
				{'O', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'O'},
			},
			expected: [][]byte{
				{'O', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'O'},
			},
		},
		{
			name: "fully surrounded O",
			board: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'O', 'X'},
				{'X', 'O', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'O', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
			},
		},
		{
			name: "single cell O",
			board: [][]byte{
				{'O'},
			},
			expected: [][]byte{
				{'O'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.board)
			for i := range tt.board {
				for j := range tt.board[i] {
					if tt.board[i][j] != tt.expected[i][j] {
						t.Errorf("at [%d][%d]: got %c, want %c", i, j, tt.board[i][j], tt.expected[i][j])
					}
				}
			}
		})
	}
}
