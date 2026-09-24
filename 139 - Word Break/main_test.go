package main

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			name:     "Example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			name:     "Single character",
			s:        "a",
			wordDict: []string{"a"},
			expected: true,
		},
		{
			name:     "Word not in dictionary",
			s:        "hello",
			wordDict: []string{"hel", "world"},
			expected: false,
		},
		{
			name:     "Repeated words allowed",
			s:        "aaab",
			wordDict: []string{"aa", "aaa", "b"},
			expected: true,
		},
		{
			name:     "No match",
			s:        "xyz",
			wordDict: []string{"abc", "def"},
			expected: false,
		},
		{
			name:     "Complete exact match",
			s:        "abc",
			wordDict: []string{"abc"},
			expected: true,
		},
		{
			name:     "Multiple valid segmentations",
			s:        "abcd",
			wordDict: []string{"a", "ab", "abc", "d"},
			expected: true,
		},
		{
			name:     "Overlap candidates",
			s:        "ccbb",
			wordDict: []string{"b", "c"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wordBreak(tt.s, tt.wordDict)
			if result != tt.expected {
				t.Errorf("wordBreak(%q, %v) = %v, want %v", tt.s, tt.wordDict, result, tt.expected)
			}
		})
	}
}
