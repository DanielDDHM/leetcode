package main

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "example 1",
			grid:     [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expected: 7,
		},
		{
			name:     "example 2",
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: 12,
		},
		{
			name:     "single cell",
			grid:     [][]int{{5}},
			expected: 5,
		},
		{
			name:     "single row",
			grid:     [][]int{{1, 2, 3, 4}},
			expected: 10,
		},
		{
			name:     "single column",
			grid:     [][]int{{1}, {2}, {3}},
			expected: 6,
		},
		{
			name:     "all zeros",
			grid:     [][]int{{0, 0}, {0, 0}},
			expected: 0,
		},
		{
			name:     "large values",
			grid:     [][]int{{200, 100}, {100, 200}},
			expected: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minPathSum(tt.grid)
			if result != tt.expected {
				t.Errorf("minPathSum(%v) = %d, want %d", tt.grid, result, tt.expected)
			}
		})
	}
}
