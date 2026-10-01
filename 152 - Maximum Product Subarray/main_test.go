package main

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{2, 3, -2, 4}, 6},
		{"example2", []int{-2, 0, -1}, 0},
		{"single element positive", []int{5}, 5},
		{"single element negative", []int{-5}, -5},
		{"all negatives", []int{-2, -3, -4}, 12},
		{"two negatives", []int{-2, 3, -4}, 24},
		{"with zero", []int{0, 2}, 2},
		{"zero at end", []int{2, 3, 0}, 6},
		{"alternating signs", []int{-1, 2, -3, 4}, 24},
		{"large negative product", []int{-2, 3, -4, 5}, 120},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProduct(tt.nums)
			if got != tt.want {
				t.Errorf("maxProduct(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
