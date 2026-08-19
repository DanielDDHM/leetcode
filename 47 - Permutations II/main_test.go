package main

import (
	"sort"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1",
			nums: []int{1, 1, 2},
			want: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name: "example 2",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "single element",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "all same",
			nums: []int{1, 1, 1},
			want: [][]int{{1, 1, 1}},
		},
		{
			name: "two duplicates",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permuteUnique(tt.nums)
			if !slicesEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func slicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	sortedA := make([][]int, len(a))
	sortedB := make([][]int, len(b))
	copy(sortedA, a)
	copy(sortedB, b)

	sort.Slice(sortedA, func(i, j int) bool {
		for k := 0; k < len(sortedA[i]); k++ {
			if sortedA[i][k] != sortedA[j][k] {
				return sortedA[i][k] < sortedA[j][k]
			}
		}
		return false
	})

	sort.Slice(sortedB, func(i, j int) bool {
		for k := 0; k < len(sortedB[i]); k++ {
			if sortedB[i][k] != sortedB[j][k] {
				return sortedB[i][k] < sortedB[j][k]
			}
		}
		return false
	})

	for i := range sortedA {
		if len(sortedA[i]) != len(sortedB[i]) {
			return false
		}
		for j := range sortedA[i] {
			if sortedA[i][j] != sortedB[i][j] {
				return false
			}
		}
	}
	return true
}
