package main

import (
	"sort"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "25525511135",
			expected: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			input:    "0000",
			expected: []string{"0.0.0.0"},
		},
		{
			input:    "101023",
			expected: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
		{
			input:    "1111",
			expected: []string{"1.1.1.1"},
		},
		{
			input:    "010010",
			expected: []string{"0.10.0.10", "0.100.1.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := restoreIpAddresses(tt.input)
			sort.Strings(result)
			sort.Strings(tt.expected)

			if len(result) != len(tt.expected) {
				t.Fatalf("expected %d results, got %d", len(tt.expected), len(result))
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("expected %s, got %s", tt.expected[i], result[i])
				}
			}
		})
	}
}
