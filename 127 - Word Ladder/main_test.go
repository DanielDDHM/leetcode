package main

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			name:      "example 1",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  5,
		},
		{
			name:      "example 2",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0,
		},
		{
			name:      "single step",
			beginWord: "a",
			endWord:   "b",
			wordList:  []string{"b"},
			expected:  2,
		},
		{
			name:      "no path",
			beginWord: "red",
			endWord:   "tax",
			wordList:  []string{"ted", "tex"},
			expected:  0,
		},
		{
			name:      "begin equals end",
			beginWord: "cat",
			endWord:   "cat",
			wordList:  []string{"cat"},
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			if got != tt.expected {
				t.Errorf("ladderLength(%q, %q, %v) = %d, want %d", tt.beginWord, tt.endWord, tt.wordList, got, tt.expected)
			}
		})
	}
}
