package main

import (
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "example1",
			input:  []int{2, 1, 5, 6, 2, 3},
			output: 10,
		},
		{
			name:   "example2",
			input:  []int{2, 4},
			output: 4,
		},
		{
			name:   "single element",
			input:  []int{5},
			output: 5,
		},
		{
			name:   "all equal elements",
			input:  []int{3, 3, 3, 3},
			output: 12,
		},
		{
			name:   "decreasing order",
			input:  []int{5, 4, 3, 2, 1},
			output: 9,
		},
		{
			name:   "increasing order",
			input:  []int{1, 2, 3, 4, 5},
			output: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestRectangleArea(tt.input)
			if result != tt.output {
				t.Errorf("largestRectangleArea(%v) = %d, want %d", tt.input, result, tt.output)
			}
		})
	}
}
