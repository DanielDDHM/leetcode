package main

import (
	"slices"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		words    []string
		expected []int
	}{
		{
			name:     "example 1",
			s:        "barfoothefoobarman",
			words:    []string{"foo", "bar"},
			expected: []int{0, 9},
		},
		{
			name:     "example 2",
			s:        "wordgoodword",
			words:    []string{"word", "good", "best"},
			expected: []int{},
		},
		{
			name:     "example 3",
			s:        "barfoobarfoobarfoobarfoobarfoo",
			words:    []string{"bar", "foo"},
			expected: []int{0, 3, 6, 9, 12, 15, 18, 21, 24},
		},
		{
			name:     "single word match",
			s:        "abc",
			words:    []string{"abc"},
			expected: []int{0},
		},
		{
			name:     "string shorter than window",
			s:        "a",
			words:    []string{"abc"},
			expected: []int{},
		},
		{
			name:     "no overlapping words",
			s:        "aabbcc",
			words:    []string{"aa", "bb"},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findSubstring(tt.s, tt.words)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
