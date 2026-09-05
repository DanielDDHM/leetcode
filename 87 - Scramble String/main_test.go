package main

import "testing"

func TestIsScramble(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"great", "rgeat", true},
		{"abcde", "caebd", false},
		{"a", "a", true},
		{"ab", "ba", true},
		{"ab", "ab", true},
		{"abc", "bca", true},
		{"abc", "acb", true},
		{"abb", "bab", true},
		{"abab", "baba", true},
		{"abab", "abab", true},
	}

	for _, tt := range tests {
		t.Run(tt.s1+","+tt.s2, func(t *testing.T) {
			result := isScramble(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("isScramble(%q, %q) = %v, expected %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}
