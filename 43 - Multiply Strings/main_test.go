package main

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		num1     string
		num2     string
		expected string
	}{
		{
			name:     "example1",
			num1:     "2",
			num2:     "3",
			expected: "6",
		},
		{
			name:     "example2",
			num1:     "123",
			num2:     "456",
			expected: "56088",
		},
		{
			name:     "zero_multiplied",
			num1:     "0",
			num2:     "123",
			expected: "0",
		},
		{
			name:     "one_multiplied",
			num1:     "1",
			num2:     "456",
			expected: "456",
		},
		{
			name:     "large_numbers",
			num1:     "999",
			num2:     "999",
			expected: "998001",
		},
		{
			name:     "single_digits",
			num1:     "9",
			num2:     "9",
			expected: "81",
		},
		{
			name:     "different_lengths",
			num1:     "12",
			num2:     "3456",
			expected: "41472",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := multiply(tt.num1, tt.num2)
			if got != tt.expected {
				t.Errorf("multiply(%q, %q) = %q, want %q", tt.num1, tt.num2, got, tt.expected)
			}
		})
	}
}
