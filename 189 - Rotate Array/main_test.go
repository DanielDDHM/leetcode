package main

import (
	"slices"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "example 2",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			name:     "single element",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "k equals length",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "k greater than length",
			nums:     []int{1, 2, 3},
			k:        5,
			expected: []int{2, 3, 1},
		},
		{
			name:     "k is zero",
			nums:     []int{1, 2, 3, 4, 5},
			k:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			if !slices.Equal(tt.nums, tt.expected) {
				t.Errorf("got %v, expected %v", tt.nums, tt.expected)
			}
		})
	}
}
