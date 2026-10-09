package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		k      int
		prices []int
		want   int
	}{
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
		{2, []int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		{1, []int{3, 2, 6, 5, 0, 3}, 4},
		{0, []int{3, 2, 6, 5, 0, 3}, 0},
		{2, []int{1}, 0},
		{2, []int{1, 2}, 1},
		{3, []int{1, 2, 3, 4, 5}, 4},
		{1, []int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := maxProfit(tt.k, tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit(%d, %v) = %d, want %d", tt.k, tt.prices, got, tt.want)
			}
		})
	}
}
