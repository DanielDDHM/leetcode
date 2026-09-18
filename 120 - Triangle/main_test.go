package main

import "testing"

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		expected int
	}{
		{
			name:     "Example 1",
			triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			expected: 11,
		},
		{
			name:     "Example 2",
			triangle: [][]int{{-10}},
			expected: -10,
		},
		{
			name:     "Single row with multiple elements",
			triangle: [][]int{{1}, {2, 3}},
			expected: 3,
		},
		{
			name:     "All negative numbers",
			triangle: [][]int{{-1}, {-2, -3}, {-4, -5, -6}},
			expected: -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumTotal(tt.triangle)
			if result != tt.expected {
				t.Errorf("minimumTotal(%v) = %d, expected %d", tt.triangle, result, tt.expected)
			}
		})
	}
}
