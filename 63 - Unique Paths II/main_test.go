package main

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]int
		expect int
	}{
		{
			name:   "Example 1",
			grid:   [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			expect: 2,
		},
		{
			name:   "Example 2",
			grid:   [][]int{{0, 1}, {0, 0}},
			expect: 1,
		},
		{
			name:   "Obstacle at start",
			grid:   [][]int{{1}},
			expect: 0,
		},
		{
			name:   "Single cell no obstacle",
			grid:   [][]int{{0}},
			expect: 1,
		},
		{
			name:   "Obstacle blocks all paths",
			grid:   [][]int{{0, 1}, {1, 0}},
			expect: 0,
		},
		{
			name:   "Vertical line",
			grid:   [][]int{{0}, {0}, {0}},
			expect: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniquePathsWithObstacles(tt.grid)
			if result != tt.expect {
				t.Errorf("got %d, want %d", result, tt.expect)
			}
		})
	}
}
