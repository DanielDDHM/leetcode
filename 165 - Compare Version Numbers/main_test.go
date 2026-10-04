package main

import (
	"testing"
)

func TestCompareVersion(t *testing.T) {
	tests := []struct {
		name     string
		version1 string
		version2 string
		expected int
	}{
		{
			name:     "Example 1",
			version1: "1.2",
			version2: "1.10",
			expected: -1,
		},
		{
			name:     "Example 2",
			version1: "1.01",
			version2: "1.001",
			expected: 0,
		},
		{
			name:     "Example 3",
			version1: "1.0",
			version2: "1.0.0.0",
			expected: 0,
		},
		{
			name:     "Version 1 greater",
			version1: "2.0",
			version2: "1.9",
			expected: 1,
		},
		{
			name:     "Single revision",
			version1: "1",
			version2: "1",
			expected: 0,
		},
		{
			name:     "Leading zeros ignored",
			version1: "1.0.1",
			version2: "1.0.01",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareVersion(tt.version1, tt.version2)
			if got != tt.expected {
				t.Errorf("compareVersion(%q, %q) = %d, want %d", tt.version1, tt.version2, got, tt.expected)
			}
		})
	}
}
