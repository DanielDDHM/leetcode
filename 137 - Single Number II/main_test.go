package main

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1",
			nums: []int{2, 2, 3, 2},
			want: 3,
		},
		{
			name: "example2",
			nums: []int{0, 1, 0, 1, 0, 1, 99},
			want: 99,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "with negative",
			nums: []int{-2, -2, -2, 1},
			want: 1,
		},
		{
			name: "negative and zero",
			nums: []int{0, 0, 0, -1},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
