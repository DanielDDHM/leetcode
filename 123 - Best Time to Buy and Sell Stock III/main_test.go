package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "example 1",
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			want:   6,
		},
		{
			name:   "example 2",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "single element",
			prices: []int{1},
			want:   0,
		},
		{
			name:   "decreasing prices",
			prices: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "two transactions with valley",
			prices: []int{1, 2, 0, 2, 4},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit(%v) = %v, want %v", tt.prices, got, tt.want)
			}
		})
	}
}
