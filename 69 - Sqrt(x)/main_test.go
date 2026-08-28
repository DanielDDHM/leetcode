package main

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"example1", 4, 2},
		{"example2", 8, 2},
		{"zero", 0, 0},
		{"one", 1, 1},
		{"two", 2, 1},
		{"fifteen", 15, 3},
		{"max", 2147483647, 46340},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mySqrt(tt.x)
			if got != tt.want {
				t.Errorf("mySqrt(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}
