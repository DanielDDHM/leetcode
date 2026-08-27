package main

import (
	"slices"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{
			name:   "example 1",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			name:   "example 2",
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2},
		},
		{
			name:   "example 3",
			digits: []int{9},
			want:   []int{1, 0},
		},
		{
			name:   "all nines",
			digits: []int{9, 9, 9},
			want:   []int{1, 0, 0, 0},
		},
		{
			name:   "zero",
			digits: []int{0},
			want:   []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if !slices.Equal(got, tt.want) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}
