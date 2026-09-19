package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"example1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"example2", []int{7, 6, 4, 3, 1}, 0},
		{"single_element", []int{1}, 0},
		{"ascending_prices", []int{1, 2, 3, 4, 5}, 4},
		{"min_at_end", []int{5, 4, 3, 2, 1}, 0},
		{"max_profit_at_end", []int{1, 5, 0, 3, 4, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
