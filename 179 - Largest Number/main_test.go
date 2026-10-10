package main

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		expect string
	}{
		{"example1", []int{10, 2}, "210"},
		{"example2", []int{3, 30, 34, 5, 9}, "9534330"},
		{"all_zeros", []int{0, 0}, "0"},
		{"single_zero", []int{0}, "0"},
		{"single_number", []int{1}, "1"},
		{"large_numbers", []int{9, 5, 34, 3, 30}, "9534330"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestNumber(tt.nums)
			if got != tt.expect {
				t.Errorf("got %q, want %q", got, tt.expect)
			}
		})
	}
}
