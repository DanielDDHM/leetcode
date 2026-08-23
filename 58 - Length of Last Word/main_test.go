package main

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example1",
			s:        "Hello World",
			expected: 5,
		},
		{
			name:     "example2",
			s:        " fly me to the moon ",
			expected: 4,
		},
		{
			name:     "example3",
			s:        "luffy is still joyboy",
			expected: 6,
		},
		{
			name:     "single_word",
			s:        "world",
			expected: 5,
		},
		{
			name:     "single_character",
			s:        "a",
			expected: 1,
		},
		{
			name:     "trailing_spaces",
			s:        "a   ",
			expected: 1,
		},
		{
			name:     "leading_spaces",
			s:        "   a",
			expected: 1,
		},
		{
			name:     "multiple_words_trailing",
			s:        "hello world   ",
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLastWord(tt.s)
			if result != tt.expected {
				t.Errorf("lengthOfLastWord(%q) = %d, expected %d", tt.s, result, tt.expected)
			}
		})
	}
}
