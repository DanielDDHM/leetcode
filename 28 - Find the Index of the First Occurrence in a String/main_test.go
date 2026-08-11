package main

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"hello", "ll", 2},
		{"a", "a", 0},
		{"abc", "c", 2},
		{"abc", "abc", 0},
		{"abc", "abcd", -1},
		{"aab", "aaab", -1},
		{"mississippi", "issip", 4},
		{"aaaa", "aaaa", 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := strStr(tt.haystack, tt.needle)
			if result != tt.expected {
				t.Errorf("strStr(%q, %q) = %d, expected %d", tt.haystack, tt.needle, result, tt.expected)
			}
		})
	}
}
