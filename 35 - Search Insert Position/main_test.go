package main

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example1",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "example2",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "example3",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			name:   "single_element_found",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "single_element_insert_after",
			nums:   []int{5},
			target: 6,
			want:   1,
		},
		{
			name:   "single_element_insert_before",
			nums:   []int{5},
			target: 4,
			want:   0,
		},
		{
			name:   "target_at_beginning",
			nums:   []int{1, 2, 3},
			target: 1,
			want:   0,
		},
		{
			name:   "target_at_end",
			nums:   []int{1, 2, 3},
			target: 3,
			want:   2,
		},
		{
			name:   "insert_before_all",
			nums:   []int{1, 2, 3},
			target: 0,
			want:   0,
		},
		{
			name:   "insert_after_all",
			nums:   []int{1, 2, 3},
			target: 4,
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.nums, tt.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
