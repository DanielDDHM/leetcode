package main

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			name:     "example 1",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:     "example 2",
			matrix:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name:     "single element",
			matrix:   [][]int{{1}},
			expected: []int{1},
		},
		{
			name:     "single row",
			matrix:   [][]int{{1, 2, 3, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "single column",
			matrix:   [][]int{{1}, {2}, {3}, {4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "negative values",
			matrix:   [][]int{{-1, -2}, {-3, -4}},
			expected: []int{-1, -2, -4, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := spiralOrder(tt.matrix)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
