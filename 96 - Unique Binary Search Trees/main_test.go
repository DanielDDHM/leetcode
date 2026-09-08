package main

import "testing"

func TestNumTrees(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"example 1", 1, 1},
		{"example 2", 3, 5},
		{"n is 2", 2, 2},
		{"n is 4", 4, 14},
		{"n is 5", 5, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numTrees(tt.n)
			if got != tt.expected {
				t.Errorf("numTrees(%d) = %d, want %d", tt.n, got, tt.expected)
			}
		})
	}
}
