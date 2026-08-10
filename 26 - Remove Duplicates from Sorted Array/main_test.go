package main

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expected      int
		expectedArray []int
	}{
		{
			name:          "example1",
			input:         []int{1, 1, 2},
			expected:      2,
			expectedArray: []int{1, 2},
		},
		{
			name:          "example2",
			input:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected:      5,
			expectedArray: []int{0, 1, 2, 3, 4},
		},
		{
			name:          "empty",
			input:         []int{},
			expected:      0,
			expectedArray: []int{},
		},
		{
			name:          "single_element",
			input:         []int{1},
			expected:      1,
			expectedArray: []int{1},
		},
		{
			name:          "all_duplicates",
			input:         []int{1, 1, 1, 1},
			expected:      1,
			expectedArray: []int{1},
		},
		{
			name:          "no_duplicates",
			input:         []int{1, 2, 3, 4, 5},
			expected:      5,
			expectedArray: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeDuplicates(tt.input)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
			if !slices.Equal(tt.input[:result], tt.expectedArray) {
				t.Errorf("got %v, want %v", tt.input[:result], tt.expectedArray)
			}
		})
	}
}
