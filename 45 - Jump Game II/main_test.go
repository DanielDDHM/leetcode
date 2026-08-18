package main

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{2, 4, 1, 1, 1, 1},
			expected: 2,
		},
		{
			name:     "example 2",
			nums:     []int{2, 1, 2, 1, 0},
			expected: 2,
		},
		{
			name:     "example 3",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "example 4",
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		{
			name:     "single element",
			nums:     []int{0},
			expected: 0,
		},
		{
			name:     "can reach end in one jump",
			nums:     []int{100, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "array of size two",
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			name:     "all zeros except first",
			nums:     []int{2, 0, 0},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := jump(tt.nums)
			if result != tt.expected {
				t.Errorf("jump(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
