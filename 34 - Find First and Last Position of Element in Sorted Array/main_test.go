package main

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "example2",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "example3",
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			name:   "single_element_target",
			nums:   []int{1},
			target: 1,
			want:   []int{0, 0},
		},
		{
			name:   "single_element_no_target",
			nums:   []int{1},
			target: 2,
			want:   []int{-1, -1},
		},
		{
			name:   "all_same_elements",
			nums:   []int{8, 8, 8},
			target: 8,
			want:   []int{0, 2},
		},
		{
			name:   "target_at_end",
			nums:   []int{1, 2, 3},
			target: 3,
			want:   []int{2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
