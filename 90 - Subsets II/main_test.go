package main

import (
	"reflect"
	"sort"
	"testing"
)

func sortSubsets(subsets [][]int) [][]int {
	sort.Slice(subsets, func(i, j int) bool {
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		for k := range subsets[i] {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})
	return subsets
}

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 2},
			expected: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 2},
				{2},
				{2, 2},
			},
		},
		{
			name: "Example 2",
			nums: []int{0},
			expected: [][]int{
				{},
				{0},
			},
		},
		{
			name: "Single element no duplicates",
			nums: []int{1},
			expected: [][]int{
				{},
				{1},
			},
		},
		{
			name: "All duplicates",
			nums: []int{2, 2, 2},
			expected: [][]int{
				{},
				{2},
				{2, 2},
				{2, 2, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subsetsWithDup(tt.nums)
			result = sortSubsets(result)
			expected := sortSubsets(tt.expected)

			if !reflect.DeepEqual(result, expected) {
				t.Errorf("got %v, want %v", result, expected)
			}
		})
	}
}
