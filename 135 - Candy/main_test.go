package main

import "testing"

func TestCandy(t *testing.T) {
	tests := []struct {
		name     string
		ratings  []int
		expected int
	}{
		{
			name:     "Example 1",
			ratings:  []int{1, 0, 2},
			expected: 5,
		},
		{
			name:     "Example 2",
			ratings:  []int{1, 2, 2},
			expected: 4,
		},
		{
			name:     "Single child",
			ratings:  []int{5},
			expected: 1,
		},
		{
			name:     "Increasing sequence",
			ratings:  []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Decreasing sequence",
			ratings:  []int{5, 4, 3, 2, 1},
			expected: 15,
		},
		{
			name:     "Valley pattern",
			ratings:  []int{1, 3, 2, 2, 1},
			expected: 7,
		},
		{
			name:     "All same ratings",
			ratings:  []int{3, 3, 3, 3},
			expected: 4,
		},
		{
			name:     "Peak in middle",
			ratings:  []int{1, 2, 3, 2, 1},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy(tt.ratings)
			if result != tt.expected {
				t.Errorf("candy(%v) = %d, expected %d", tt.ratings, result, tt.expected)
			}
		})
	}
}
