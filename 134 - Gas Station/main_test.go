package main

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas      []int
		cost     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
		{[]int{1}, []int{1}, 0},
		{[]int{0}, []int{1}, -1},
		{[]int{5, 5, 5}, []int{5, 5, 5}, 0},
		{[]int{1, 1, 1}, []int{2, 2, 2}, -1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := canCompleteCircuit(tt.gas, tt.cost)
			if result != tt.expected {
				t.Errorf("canCompleteCircuit(%v, %v) = %d, expected %d", tt.gas, tt.cost, result, tt.expected)
			}
		})
	}
}
