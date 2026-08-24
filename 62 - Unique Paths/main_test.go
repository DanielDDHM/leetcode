package main

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{"Example 1: 3x7 grid", 3, 7, 28},
		{"Example 2: 3x2 grid", 3, 2, 3},
		{"Edge case: 1x1 grid", 1, 1, 1},
		{"Edge case: 1x10 grid", 1, 10, 1},
		{"Edge case: 10x1 grid", 10, 1, 1},
		{"Edge case: 2x2 grid", 2, 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniquePaths(tt.m, tt.n)
			if result != tt.expected {
				t.Errorf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, result, tt.expected)
			}
		})
	}
}
