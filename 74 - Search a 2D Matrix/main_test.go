package main

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			name:     "Example 1: target found in middle",
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
		},
		{
			name:     "Example 2: target not found",
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   13,
			expected: false,
		},
		{
			name:     "Single element matrix: target present",
			matrix:   [][]int{{5}},
			target:   5,
			expected: true,
		},
		{
			name:     "Single element matrix: target not present",
			matrix:   [][]int{{5}},
			target:   3,
			expected: false,
		},
		{
			name:     "Target at start",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			target:   1,
			expected: true,
		},
		{
			name:     "Target at end",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			target:   9,
			expected: true,
		},
		{
			name:     "Negative numbers",
			matrix:   [][]int{{-10, -8, -6, -4}, {-3, -1, 0, 2}, {3, 5, 7, 9}},
			target:   -6,
			expected: true,
		},
		{
			name:     "Target smaller than all elements",
			matrix:   [][]int{{5, 10, 15}, {20, 25, 30}},
			target:   1,
			expected: false,
		},
		{
			name:     "Target larger than all elements",
			matrix:   [][]int{{5, 10, 15}, {20, 25, 30}},
			target:   100,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchMatrix(tt.matrix, tt.target)
			if result != tt.expected {
				t.Errorf("searchMatrix(%v, %d) = %v, want %v", tt.matrix, tt.target, result, tt.expected)
			}
		})
	}
}
