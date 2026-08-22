package main

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "Example 1",
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:     "Example 2",
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "Example 3",
			input:    [][]int{{4, 7}, {1, 4}},
			expected: [][]int{{1, 7}},
		},
		{
			name:     "Single interval",
			input:    [][]int{{1, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "No overlapping intervals",
			input:    [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := merge(tt.input)
			if !slices.EqualFunc(result, tt.expected, func(a, b []int) bool {
				return len(a) == len(b) && a[0] == b[0] && a[1] == b[1]
			}) {
				t.Errorf("merge(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
