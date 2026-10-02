package main

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1_not_rotated",
			nums: []int{1, 3, 5},
			want: 1,
		},
		{
			name: "example2_with_duplicates",
			nums: []int{2, 2, 2, 0, 1},
			want: 0,
		},
		{
			name: "example3_all_same",
			nums: []int{3, 3, 3, 3, 3},
			want: 3,
		},
		{
			name: "single_element",
			nums: []int{5},
			want: 5,
		},
		{
			name: "two_elements_ascending",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "two_elements_rotated",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "rotated_no_duplicates",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "duplicates_at_boundaries",
			nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1},
			want: 1,
		},
		{
			name: "duplicates_at_rotation_point",
			nums: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
