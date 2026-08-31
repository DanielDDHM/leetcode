package main

import (
	"slices"
	"sort"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want [][]int
	}{
		{
			name: "example 1",
			n:    4,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "example 2",
			n:    1,
			k:    1,
			want: [][]int{{1}},
		},
		{
			name: "n=2 k=1",
			n:    2,
			k:    1,
			want: [][]int{{1}, {2}},
		},
		{
			name: "n=3 k=3",
			n:    3,
			k:    3,
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "n=5 k=3",
			n:    5,
			k:    3,
			want: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combine(tt.n, tt.k)

			if len(got) != len(tt.want) {
				t.Errorf("got %d combinations, want %d", len(got), len(tt.want))
				return
			}

			sort.Slice(got, func(i, j int) bool {
				for k := range got[i] {
					if got[i][k] != got[j][k] {
						return got[i][k] < got[j][k]
					}
				}
				return false
			})
			sort.Slice(tt.want, func(i, j int) bool {
				for k := range tt.want[i] {
					if tt.want[i][k] != tt.want[j][k] {
						return tt.want[i][k] < tt.want[j][k]
					}
				}
				return false
			})

			for i := range got {
				if !slices.Equal(got[i], tt.want[i]) {
					t.Errorf("combination %d: got %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
