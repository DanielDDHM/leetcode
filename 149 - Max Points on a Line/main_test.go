package main

import "testing"

func TestMaxPoints(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name:     "Example 1",
			points:   [][]int{{1, 1}, {2, 2}, {3, 3}},
			expected: 3,
		},
		{
			name:     "Example 2",
			points:   [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
			expected: 4,
		},
		{
			name:     "Single point",
			points:   [][]int{{0, 0}},
			expected: 1,
		},
		{
			name:     "Two points",
			points:   [][]int{{0, 0}, {1, 1}},
			expected: 2,
		},
		{
			name:     "All same point",
			points:   [][]int{{0, 0}, {0, 0}, {0, 0}},
			expected: 3,
		},
		{
			name:     "Vertical line",
			points:   [][]int{{0, 0}, {0, 1}, {0, 2}},
			expected: 3,
		},
		{
			name:     "Horizontal line",
			points:   [][]int{{0, 0}, {1, 0}, {2, 0}},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxPoints(tt.points)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
