package main

import (
	"slices"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "example1",
			input:  []int{2, 0, 2, 1, 1, 0},
			expect: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:   "example2",
			input:  []int{2, 0, 1},
			expect: []int{0, 1, 2},
		},
		{
			name:   "all_zeros",
			input:  []int{0, 0, 0},
			expect: []int{0, 0, 0},
		},
		{
			name:   "all_twos",
			input:  []int{2, 2, 2},
			expect: []int{2, 2, 2},
		},
		{
			name:   "single_element",
			input:  []int{1},
			expect: []int{1},
		},
		{
			name:   "reverse_sorted",
			input:  []int{2, 2, 2, 1, 1, 1, 0, 0, 0},
			expect: []int{0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.input)
			if !slices.Equal(tt.input, tt.expect) {
				t.Errorf("got %v, expect %v", tt.input, tt.expect)
			}
		})
	}
}
