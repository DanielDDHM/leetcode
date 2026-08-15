package main

import "testing"

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected string
	}{
		{"n=1", 1, "1"},
		{"n=2", 2, "11"},
		{"n=3", 3, "21"},
		{"n=4", 4, "1211"},
		{"n=5", 5, "111221"},
		{"n=6", 6, "312211"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countAndSay(tt.n)
			if result != tt.expected {
				t.Errorf("countAndSay(%d) = %q, want %q", tt.n, result, tt.expected)
			}
		})
	}
}
