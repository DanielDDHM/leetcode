package main

import "testing"

func TestMinCut(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1: aab",
			s:        "aab",
			expected: 1,
		},
		{
			name:     "Example 2: a",
			s:        "a",
			expected: 0,
		},
		{
			name:     "Example 3: ab",
			s:        "ab",
			expected: 1,
		},
		{
			name:     "Palindrome: aba",
			s:        "aba",
			expected: 0,
		},
		{
			name:     "All different: abcd",
			s:        "abcd",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minCut(tt.s)
			if result != tt.expected {
				t.Errorf("minCut(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
