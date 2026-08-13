package main

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "example2",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "single_element_found",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "single_element_not_found",
			nums:   []int{1},
			target: 3,
			want:   -1,
		},
		{
			name:   "target_at_start",
			nums:   []int{3, 1},
			target: 3,
			want:   0,
		},
		{
			name:   "target_at_end",
			nums:   []int{1, 3},
			target: 3,
			want:   1,
		},
		{
			name:   "no_rotation",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "rotated_pivot_1",
			nums:   []int{5, 1, 2, 3, 4},
			target: 1,
			want:   1,
		},
		{
			name:   "rotated_pivot_4",
			nums:   []int{2, 3, 4, 5, 1},
			target: 1,
			want:   4,
		},
		{
			name:   "target_left_side",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 6,
			want:   2,
		},
		{
			name:   "target_right_side",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 1,
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
