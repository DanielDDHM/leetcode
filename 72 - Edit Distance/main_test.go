package main

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{"example1", "horse", "ros", 3},
		{"example2", "intention", "execution", 5},
		{"empty_both", "", "", 0},
		{"empty_word1", "", "abc", 3},
		{"empty_word2", "abc", "", 3},
		{"identical", "test", "test", 0},
		{"single_char_same", "a", "a", 0},
		{"single_char_different", "a", "b", 1},
		{"complete_replace", "abc", "xyz", 3},
		{"insert_only", "a", "abc", 2},
		{"delete_only", "abc", "a", 2},
		{"mixed_operations", "ab", "ba", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistance(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("minDistance(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}
