package main

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{
			name:    "Example 1",
			numRows: 5,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			name:    "Example 2",
			numRows: 1,
			want: [][]int{
				{1},
			},
		},
		{
			name:    "Edge case: 2 rows",
			numRows: 2,
			want: [][]int{
				{1},
				{1, 1},
			},
		},
		{
			name:    "Edge case: 3 rows",
			numRows: 3,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generate(tt.numRows)
			if len(got) != len(tt.want) {
				t.Fatalf("generate(%d) returned %d rows, want %d rows", tt.numRows, len(got), len(tt.want))
			}
			for i := range got {
				if len(got[i]) != len(tt.want[i]) {
					t.Fatalf("row %d has %d elements, want %d", i, len(got[i]), len(tt.want[i]))
				}
				for j := range got[i] {
					if got[i][j] != tt.want[i][j] {
						t.Errorf("row %d, col %d: got %d, want %d", i, j, got[i][j], tt.want[i][j])
					}
				}
			}
		})
	}
}
