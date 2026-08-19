package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name: "4x4 matrix",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			want: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			name:   "1x1 matrix",
			matrix: [][]int{{1}},
			want:   [][]int{{1}},
		},
		{
			name: "2x2 matrix",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			want: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			name: "3x3 matrix",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.matrix)
			for i := 0; i < len(tt.matrix); i++ {
				for j := 0; j < len(tt.matrix[i]); j++ {
					if tt.matrix[i][j] != tt.want[i][j] {
						t.Errorf("matrix[%d][%d] = %d, want %d", i, j, tt.matrix[i][j], tt.want[i][j])
					}
				}
			}
		})
	}
}
