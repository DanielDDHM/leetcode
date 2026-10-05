package main

import (
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	tests := []struct {
		name        string
		numerator   int
		denominator int
		expected    string
	}{
		{
			name:        "Example 1",
			numerator:   1,
			denominator: 2,
			expected:    "0.5",
		},
		{
			name:        "Example 2",
			numerator:   2,
			denominator: 1,
			expected:    "2",
		},
		{
			name:        "Example 3",
			numerator:   4,
			denominator: 333,
			expected:    "0.(012)",
		},
		{
			name:        "Negative numerator",
			numerator:   -1,
			denominator: 2,
			expected:    "-0.5",
		},
		{
			name:        "Negative denominator",
			numerator:   1,
			denominator: -2,
			expected:    "-0.5",
		},
		{
			name:        "Both negative",
			numerator:   -1,
			denominator: -2,
			expected:    "0.5",
		},
		{
			name:        "Zero numerator",
			numerator:   0,
			denominator: 1,
			expected:    "0",
		},
		{
			name:        "Repeating 1/6",
			numerator:   1,
			denominator: 6,
			expected:    "0.1(6)",
		},
		{
			name:        "Integer result",
			numerator:   6,
			denominator: 2,
			expected:    "3",
		},
		{
			name:        "Complex repeating",
			numerator:   1,
			denominator: 3,
			expected:    "0.(3)",
		},
		{
			name:        "Negative with repeating",
			numerator:   -1,
			denominator: 3,
			expected:    "-0.(3)",
		},
		{
			name:        "Large numerator",
			numerator:   2147483647,
			denominator: 1,
			expected:    "2147483647",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fractionToDecimal(tt.numerator, tt.denominator)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
