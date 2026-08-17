package main

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s        string
		p        string
		expected bool
	}{
		{s: "aa", p: "a", expected: false},
		{s: "aa", p: "*", expected: true},
		{s: "cb", p: "?a", expected: false},
		{s: "adceb", p: "*a*b", expected: true},
		{s: "acdcb", p: "a*c?b", expected: false},
		{s: "", p: "", expected: true},
		{s: "", p: "*", expected: true},
		{s: "a", p: "", expected: false},
		{s: "a", p: "*", expected: true},
		{s: "a", p: "?", expected: true},
		{s: "a", p: "a", expected: true},
		{s: "a", p: "b", expected: false},
		{s: "aa", p: "?", expected: false},
		{s: "aa", p: "??", expected: true},
		{s: "ab", p: "a*", expected: true},
		{s: "ab", p: "*b", expected: true},
		{s: "ab", p: "*c", expected: false},
		{s: "abc", p: "a*c", expected: true},
		{s: "abc", p: "a*d", expected: false},
		{s: "mississippi", p: "m*iss*p*..", expected: false},
		{s: "aa", p: "*?", expected: true},
		{s: "a", p: "*?", expected: true},
		{s: "", p: "***", expected: true},
		{s: "b", p: "*?", expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.s+"_"+tt.p, func(t *testing.T) {
			result := isMatch(tt.s, tt.p)
			if result != tt.expected {
				t.Errorf("isMatch(%q, %q) = %v, want %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}
