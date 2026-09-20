package main

import (
	"sort"
	"testing"
)

func TestFindLadders(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		expected  [][]string
	}{
		{
			name:      "example 1",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected: [][]string{
				{"hit", "hot", "dot", "dog", "cog"},
				{"hit", "hot", "lot", "log", "cog"},
			},
		},
		{
			name:      "example 2",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  [][]string{},
		},
		{
			name:      "single word match",
			beginWord: "a",
			endWord:   "b",
			wordList:  []string{"b"},
			expected: [][]string{
				{"a", "b"},
			},
		},
		{
			name:      "no path",
			beginWord: "red",
			endWord:   "tax",
			wordList:  []string{"ted", "tex"},
			expected:  [][]string{},
		},
		{
			name:      "end word not in list",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findLadders(tt.beginWord, tt.endWord, tt.wordList)
			if !slicesEqual(sortLadders(result), sortLadders(tt.expected)) {
				t.Errorf("findLadders(%q, %q, %v) = %v, want %v", tt.beginWord, tt.endWord, tt.wordList, result, tt.expected)
			}
		})
	}
}

func slicesEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !sliceStringEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func sliceStringEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sortLadders(ladders [][]string) [][]string {
	result := make([][]string, len(ladders))
	for i, ladder := range ladders {
		result[i] = make([]string, len(ladder))
		copy(result[i], ladder)
	}
	sort.Slice(result, func(i, j int) bool {
		for k := range result[i] {
			if result[i][k] < result[j][k] {
				return true
			}
			if result[i][k] > result[j][k] {
				return false
			}
		}
		return false
	})
	return result
}
