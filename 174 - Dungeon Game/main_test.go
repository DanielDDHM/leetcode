package main

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	tests := []struct {
		name     string
		dungeon  [][]int
		expected int
	}{
		{
			name:     "Example 1",
			dungeon:  [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}},
			expected: 7,
		},
		{
			name:     "Example 2",
			dungeon:  [][]int{{0}},
			expected: 1,
		},
		{
			name:     "Single cell negative",
			dungeon:  [][]int{{-3}},
			expected: 4,
		},
		{
			name:     "Single cell positive",
			dungeon:  [][]int{{5}},
			expected: 1,
		},
		{
			name:     "All positive values",
			dungeon:  [][]int{{1, 2}, {3, 4}},
			expected: 1,
		},
		{
			name:     "2x2 with negatives",
			dungeon:  [][]int{{-2, 1}, {1, -1}},
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calculateMinimumHP(test.dungeon)
			if result != test.expected {
				t.Errorf("got %d, want %d", result, test.expected)
			}
		})
	}
}
