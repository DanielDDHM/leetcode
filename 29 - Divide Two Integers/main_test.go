package main

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		dividend int
		divisor  int
		expected int
	}{
		{"Example 1", 10, 3, 3},
		{"Example 2", 7, -3, -2},
		{"Overflow case", -2147483648, -1, 2147483647},
		{"Min value with 1", -2147483648, 1, -2147483648},
		{"Zero dividend", 0, 1, 0},
		{"Equal operands", 1, 1, 1},
		{"Negative dividend", -10, 3, -3},
		{"Large positive", 2147483647, 1, 2147483647},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := divide(tt.dividend, tt.divisor)
			if result != tt.expected {
				t.Errorf("divide(%d, %d) = %d, expected %d", tt.dividend, tt.divisor, result, tt.expected)
			}
		})
	}
}
