package main

import (
	"fmt"
	"testing"
)

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{4, 2},
		{8, 92},
		{9, 352},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			got := totalNQueens(tt.n)
			if got != tt.want {
				t.Errorf("totalNQueens(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
