package main

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "example 1",
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			name:     "example 2",
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
		{
			name:     "both zero",
			a:        "0",
			b:        "0",
			expected: "0",
		},
		{
			name:     "simple carry",
			a:        "1",
			b:        "1",
			expected: "10",
		},
		{
			name:     "multiple carries",
			a:        "1111",
			b:        "1111",
			expected: "11110",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addBinary(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("addBinary(%q, %q) = %q, want %q", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
