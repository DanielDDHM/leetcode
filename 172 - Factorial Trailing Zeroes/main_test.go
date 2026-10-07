package main

import "testing"

func TestTrailingZeroes(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "example1",
			n:        3,
			expected: 0,
		},
		{
			name:     "example2",
			n:        5,
			expected: 1,
		},
		{
			name:     "example3",
			n:        0,
			expected: 0,
		},
		{
			name:     "divisible_by_25",
			n:        25,
			expected: 6,
		},
		{
			name:     "divisible_by_125",
			n:        125,
			expected: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trailingZeroes(tt.n)
			if result != tt.expected {
				t.Errorf("trailingZeroes(%d) = %d, want %d", tt.n, result, tt.expected)
			}
		})
	}
}
