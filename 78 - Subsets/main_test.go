package main

import (
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example1",
			nums: []int{1, 2, 3},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			name: "example2",
			nums: []int{0},
			want: [][]int{
				{},
				{0},
			},
		},
		{
			name: "single positive",
			nums: []int{5},
			want: [][]int{
				{},
				{5},
			},
		},
		{
			name: "two elements",
			nums: []int{1, 2},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
			},
		},
		{
			name: "negative numbers",
			nums: []int{-1, 0, 1},
			want: [][]int{
				{},
				{-1},
				{0},
				{-1, 0},
				{1},
				{-1, 1},
				{0, 1},
				{-1, 0, 1},
			},
		},
		{
			name: "four elements",
			nums: []int{1, 2, 3, 4},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
				{4},
				{1, 4},
				{2, 4},
				{1, 2, 4},
				{3, 4},
				{1, 3, 4},
				{2, 3, 4},
				{1, 2, 3, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)

			if len(got) != len(tt.want) {
				t.Errorf("length mismatch: got %d, want %d", len(got), len(tt.want))
				return
			}

			gotSorted := make([][]int, len(got))
			wantSorted := make([][]int, len(tt.want))

			for i := range got {
				gotSorted[i] = make([]int, len(got[i]))
				copy(gotSorted[i], got[i])
				sort.Ints(gotSorted[i])
			}

			for i := range tt.want {
				wantSorted[i] = make([]int, len(tt.want[i]))
				copy(wantSorted[i], tt.want[i])
				sort.Ints(wantSorted[i])
			}

			sort.Slice(gotSorted, func(i, j int) bool {
				for k := 0; k < len(gotSorted[i]) && k < len(gotSorted[j]); k++ {
					if gotSorted[i][k] != gotSorted[j][k] {
						return gotSorted[i][k] < gotSorted[j][k]
					}
				}
				return len(gotSorted[i]) < len(gotSorted[j])
			})

			sort.Slice(wantSorted, func(i, j int) bool {
				for k := 0; k < len(wantSorted[i]) && k < len(wantSorted[j]); k++ {
					if wantSorted[i][k] != wantSorted[j][k] {
						return wantSorted[i][k] < wantSorted[j][k]
					}
				}
				return len(wantSorted[i]) < len(wantSorted[j])
			})

			for i := range gotSorted {
				if len(gotSorted[i]) != len(wantSorted[i]) {
					t.Errorf("subset %d length mismatch: got %v, want %v", i, gotSorted[i], wantSorted[i])
					continue
				}
				for j := range gotSorted[i] {
					if gotSorted[i][j] != wantSorted[i][j] {
						t.Errorf("subset %d mismatch: got %v, want %v", i, gotSorted[i], wantSorted[i])
						break
					}
				}
			}
		})
	}
}
