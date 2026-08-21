package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "example 2",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "example 3",
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
		{
			name:     "all negative",
			nums:     []int{-5, -3, -2},
			expected: -2,
		},
		{
			name:     "single negative",
			nums:     []int{-100},
			expected: -100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSubArray(tt.nums)
			if result != tt.expected {
				t.Errorf("maxSubArray(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
