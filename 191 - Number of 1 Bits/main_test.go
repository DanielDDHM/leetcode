package main

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"Example 1", 11, 3},
		{"Example 2", 128, 1},
		{"Example 3", 2147483645, 30},
		{"Edge case 1", 1, 1},
		{"Edge case 2", 2147483647, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.n); got != tt.want {
				t.Errorf("hammingWeight(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
