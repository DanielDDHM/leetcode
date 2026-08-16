package main

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "example 1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "example 2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "no water trapped",
			height: []int{1, 2, 3, 4, 5},
			want:   0,
		},
		{
			name:   "single bar",
			height: []int{1},
			want:   0,
		},
		{
			name:   "valley between two bars",
			height: []int{5, 0, 0, 0, 5},
			want:   15,
		},
		{
			name:   "asymmetric valley",
			height: []int{0, 2, 0, 4, 0, 3, 0},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
