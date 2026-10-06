package main

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example1",
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "example2",
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "negative single",
			nums:     []int{-1},
			expected: -1,
		},
		{
			name:     "large array",
			nums:     []int{1, 1, 1, 1, 0, 0, 0},
			expected: 1,
		},
		{
			name:     "min int value",
			nums:     []int{-2147483648},
			expected: -2147483648,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := majorityElement(tt.nums)
			if result != tt.expected {
				t.Errorf("majorityElement(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
