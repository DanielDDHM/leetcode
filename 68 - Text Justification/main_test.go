package main

import (
	"testing"
)

func TestFullJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		expected []string
	}{
		{
			name:     "example 1",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification"},
			maxWidth: 16,
			expected: []string{"This    is    an", "example  of text", "justification   "},
		},
		{
			name:     "example 2",
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			expected: []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			name:     "example 3",
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer", "while", "art", "is", "everything", "else"},
			maxWidth: 20,
			expected: []string{"Science  is  what we", "understand      well", "enough to explain to", "a computer while art", "is everything else  "},
		},
		{
			name:     "single word",
			words:    []string{"justification"},
			maxWidth: 16,
			expected: []string{"justification   "},
		},
		{
			name:     "exact fit",
			words:    []string{"a", "b", "c"},
			maxWidth: 5,
			expected: []string{"a b c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fullJustify(tt.words, tt.maxWidth)
			if len(result) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, want %d", len(result), len(tt.expected))
			}
			for i := range result {
				if i < len(tt.expected) {
					if result[i] != tt.expected[i] {
						t.Errorf("line %d mismatch: got %q, want %q", i, result[i], tt.expected[i])
					}
					if len(result[i]) != tt.maxWidth {
						t.Errorf("line %d length mismatch: got %d, want %d", i, len(result[i]), tt.maxWidth)
					}
				}
			}
		})
	}
}
