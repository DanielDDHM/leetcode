package main

import (
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 0},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 4, -1, 1},
			expected: 2,
		},
		{
			name:     "Example 3",
			nums:     []int{7, 8, 9, 11, 12},
			expected: 1,
		},
		{
			name:     "Single element 1",
			nums:     []int{1},
			expected: 2,
		},
		{
			name:     "Single element not 1",
			nums:     []int{2},
			expected: 1,
		},
		{
			name:     "Consecutive from 1",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "Large gap",
			nums:     []int{1000000},
			expected: 1,
		},
		{
			name:     "All negative",
			nums:     []int{-1, -2, -3},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			result := firstMissingPositive(numsCopy)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
