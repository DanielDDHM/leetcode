package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "example1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			name:   "example2",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "single_element",
			prices: []int{1},
			want:   0,
		},
		{
			name:   "descending_prices",
			prices: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "constant_prices",
			prices: []int{1, 1, 1, 1},
			want:   0,
		},
		{
			name:   "multiple_transactions",
			prices: []int{1, 4, 2, 7},
			want:   8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
