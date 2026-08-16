package main

import (
	"slices"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "example_1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			name:       "example_2",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected: [][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			name:       "single_element",
			candidates: []int{5},
			target:     5,
			expected: [][]int{
				{5},
			},
		},
		{
			name:       "all_equal_elements",
			candidates: []int{1, 1, 1, 1, 1},
			target:     3,
			expected: [][]int{
				{1, 1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combinationSum2(tt.candidates, tt.target)

			if len(result) != len(tt.expected) {
				t.Errorf("got %d combinations, expected %d", len(result), len(tt.expected))
				return
			}

			for _, comb := range result {
				slices.Sort(comb)
			}
			for _, exp := range tt.expected {
				slices.Sort(exp)
			}

			slices.SortFunc(result, func(a, b []int) int {
				for i := 0; i < len(a) && i < len(b); i++ {
					if a[i] < b[i] {
						return -1
					}
					if a[i] > b[i] {
						return 1
					}
				}
				return len(a) - len(b)
			})
			slices.SortFunc(tt.expected, func(a, b []int) int {
				for i := 0; i < len(a) && i < len(b); i++ {
					if a[i] < b[i] {
						return -1
					}
					if a[i] > b[i] {
						return 1
					}
				}
				return len(a) - len(b)
			})

			for i, comb := range result {
				if !slices.Equal(comb, tt.expected[i]) {
					t.Errorf("combination %d: got %v, expected %v", i, comb, tt.expected[i])
				}
			}
		})
	}
}
