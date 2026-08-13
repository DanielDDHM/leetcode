package main

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example1",
			input:    "(()",
			expected: 2,
		},
		{
			name:     "example2",
			input:    ")()())",
			expected: 4,
		},
		{
			name:     "empty",
			input:    "",
			expected: 0,
		},
		{
			name:     "single_open",
			input:    "(",
			expected: 0,
		},
		{
			name:     "single_close",
			input:    ")",
			expected: 0,
		},
		{
			name:     "simple_pair",
			input:    "()",
			expected: 2,
		},
		{
			name:     "multiple_pairs",
			input:    "()()",
			expected: 4,
		},
		{
			name:     "nested_valid",
			input:    "(())",
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestValidParentheses(tt.input)
			if result != tt.expected {
				t.Errorf("longestValidParentheses(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
