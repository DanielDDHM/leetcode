package main

import "testing"

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"Example 1", 1, "A"},
		{"Example 2", 28, "AB"},
		{"Example 3", 701, "ZY"},
		{"Single letter Z", 26, "Z"},
		{"Max int32", 2147483647, "FXSHRXW"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToTitle(tt.input)
			if result != tt.expected {
				t.Errorf("convertToTitle(%d) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}
