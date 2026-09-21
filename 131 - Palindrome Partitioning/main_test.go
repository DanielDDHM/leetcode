package main

import (
	"slices"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [][]string
	}{
		{
			name:  "example1",
			input: "aab",
			expected: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name:  "example2",
			input: "a",
			expected: [][]string{
				{"a"},
			},
		},
		{
			name:  "all_same",
			input: "aaa",
			expected: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name:  "no_adjacent_palindromes",
			input: "abc",
			expected: [][]string{
				{"a", "b", "c"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partition(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("got %d partitions, want %d", len(result), len(tt.expected))
				return
			}

			for _, exp := range tt.expected {
				found := false
				for _, res := range result {
					if slices.Equal(res, exp) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected partition %v not found in result %v", exp, result)
				}
			}
		})
	}
}
