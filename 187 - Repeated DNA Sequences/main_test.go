package main

import (
	"slices"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []string
	}{
		{
			name:     "Example 1",
			s:        "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			expected: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			name:     "Example 2",
			s:        "AAAAAAAAAAAAA",
			expected: []string{"AAAAAAAAAA"},
		},
		{
			name:     "No repeated sequences",
			s:        "ACGTACGTAAACGT",
			expected: []string{},
		},
		{
			name:     "Less than 10 characters",
			s:        "ACGTACGT",
			expected: []string{},
		},
		{
			name:     "Exactly 10 characters no repeat",
			s:        "ACGTACGTAA",
			expected: []string{},
		},
		{
			name:     "Single repeated sequence",
			s:        "AAAAAAAAAA" + "GGGGGGGGGG" + "AAAAAAAAAA",
			expected: []string{"AAAAAAAAAA"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findRepeatedDnaSequences(tt.s)
			slices.Sort(result)
			slices.Sort(tt.expected)
			if len(result) != len(tt.expected) {
				t.Errorf("Length mismatch: got %d, expected %d", len(result), len(tt.expected))
				return
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("Mismatch at index %d: got %q, expected %q", i, v, tt.expected[i])
				}
			}
		})
	}
}
