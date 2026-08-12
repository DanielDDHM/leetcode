package main

import (
	"slices"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "example 1",
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "example 2",
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "example 3",
			input:    []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "two elements ascending",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "two elements descending",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "all same",
			input:    []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		{
			name:     "last permutation",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.input)
			if !slices.Equal(tt.input, tt.expected) {
				t.Errorf("got %v, want %v", tt.input, tt.expected)
			}
		})
	}
}
