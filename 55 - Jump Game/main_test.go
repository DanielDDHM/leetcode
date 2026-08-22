package main

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"Example 1", []int{2, 3, 1, 1, 4}, true},
		{"Example 2", []int{3, 2, 1, 0, 4}, false},
		{"Single element", []int{0}, true},
		{"Can reach end with jump", []int{2, 0, 0}, true},
		{"Stuck at first", []int{0, 1}, false},
		{"One element non-zero", []int{1}, true},
		{"Large jump from start", []int{5, 1, 1, 1, 1}, true},
		{"Zero after non-zero", []int{1, 0, 1}, false},
		{"All ones", []int{1, 1, 1, 1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.nums); got != tt.want {
				t.Errorf("canJump(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
