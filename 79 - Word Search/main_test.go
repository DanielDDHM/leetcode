package main

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name: "Example 1 - ABCCED exists",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
		{
			name: "Example 2 - SEE exists",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			want: true,
		},
		{
			name: "Example 3 - ABCB does not exist",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
			want: false,
		},
		{
			name:  "Single cell match",
			board: [][]byte{{'A'}},
			word:  "A",
			want:  true,
		},
		{
			name: "Word not in board",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word: "XYZ",
			want: false,
		},
		{
			name: "Word longer than path available",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word: "ABCDEFGH",
			want: false,
		},
		{
			name: "Requires reusing cell - not allowed",
			board: [][]byte{
				{'A', 'A'},
				{'A', 'A'},
			},
			word: "AAAA",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exist(tt.board, tt.word)
			if got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
