package main

import (
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "example 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			name:       "example 2",
			candidates: []int{2, 3, 5},
			target:     8,
			expected: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			name:       "example 3",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "single candidate match",
			candidates: []int{5},
			target:     5,
			expected: [][]int{
				{5},
			},
		},
		{
			name:       "multiple of same element",
			candidates: []int{1, 2},
			target:     3,
			expected: [][]int{
				{1, 1, 1},
				{1, 2},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := combinationSum(test.candidates, test.target)
			if !slicesEqual(result, test.expected) {
				t.Errorf("got %v, want %v", result, test.expected)
			}
		})
	}
}

func slicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, subA := range a {
		sort.Ints(subA)
	}
	for _, subB := range b {
		sort.Ints(subB)
	}

	sort.Slice(a, func(i, j int) bool {
		for k := 0; k < len(a[i]) && k < len(a[j]); k++ {
			if a[i][k] != a[j][k] {
				return a[i][k] < a[j][k]
			}
		}
		return len(a[i]) < len(a[j])
	})

	sort.Slice(b, func(i, j int) bool {
		for k := 0; k < len(b[i]) && k < len(b[j]); k++ {
			if b[i][k] != b[j][k] {
				return b[i][k] < b[j][k]
			}
		}
		return len(b[i]) < len(b[j])
	})

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
