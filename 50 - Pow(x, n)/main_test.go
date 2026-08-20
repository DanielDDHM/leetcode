package main

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		n        int
		expected float64
	}{
		{"example1", 2.0, 10, 1024.0},
		{"example2", 2.1, 3, 9.261},
		{"example3", 2.0, -2, 0.25},
		{"edge_case_one", 1.0, 2147483647, 1.0},
		{"edge_case_negative_base_even", -1.0, 2, 1.0},
		{"edge_case_negative_base_odd", -1.0, 3, -1.0},
		{"edge_case_zero", 0.0, 1, 0.0},
		{"edge_case_min_int", 1.0, -2147483648, 1.0},
		{"edge_case_zero_exponent", 2.0, 0, 1.0},
		{"edge_case_negative_exponent", 0.5, -2, 4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := myPow(tt.x, tt.n)
			if math.Abs(result-tt.expected) > 1e-6 {
				t.Errorf("myPow(%v, %d) = %v, want %v", tt.x, tt.n, result, tt.expected)
			}
		})
	}
}
