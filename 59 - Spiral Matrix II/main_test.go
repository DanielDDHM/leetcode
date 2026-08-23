package main

import (
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected [][]int
	}{
		{
			name: "example 1",
			n:    3,
			expected: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "example 2",
			n:    1,
			expected: [][]int{
				{1},
			},
		},
		{
			name: "n equals 2",
			n:    2,
			expected: [][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			name: "n equals 4",
			n:    4,
			expected: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateMatrix(tt.n)
			if !matricesEqual(result, tt.expected) {
				t.Errorf("generateMatrix(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}

func matricesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
