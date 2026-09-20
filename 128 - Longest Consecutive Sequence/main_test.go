package main

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Example 1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"Empty array", []int{}, 0},
		{"Single element", []int{1}, 1},
		{"All duplicates", []int{1, 1, 1, 1}, 1},
		{"Negative numbers", []int{-1, 0, 1}, 3},
		{"No consecutive", []int{10, 20, 30}, 1},
		{"All consecutive", []int{5, 2, 4, 3, 1}, 5},
		{"Large gap", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestConsecutive(tt.nums)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
