package main

import (
	"sort"
	"testing"
)

func sortResults(results []string) []string {
	sorted := make([]string, len(results))
	copy(sorted, results)
	sort.Strings(sorted)
	return sorted
}

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     []string
	}{
		{
			name:     "Example 1",
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			want:     []string{"cats and dog", "cat sand dog"},
		},
		{
			name:     "Example 2",
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			want:     []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			name:     "Example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     []string{},
		},
		{
			name:     "Single word",
			s:        "a",
			wordDict: []string{"a"},
			want:     []string{"a"},
		},
		{
			name:     "Impossible",
			s:        "a",
			wordDict: []string{"b"},
			want:     []string{},
		},
		{
			name:     "Multiple segments",
			s:        "aaa",
			wordDict: []string{"a", "aa"},
			want:     []string{"a a a", "aa a", "a aa"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.s, tt.wordDict)
			gotSorted := sortResults(got)
			wantSorted := sortResults(tt.want)

			if len(gotSorted) != len(wantSorted) {
				t.Errorf("got %d results, want %d", len(gotSorted), len(wantSorted))
				t.Errorf("got: %v", gotSorted)
				t.Errorf("want: %v", wantSorted)
				return
			}

			for i := range gotSorted {
				if gotSorted[i] != wantSorted[i] {
					t.Errorf("result mismatch at index %d", i)
					t.Errorf("got: %v", gotSorted)
					t.Errorf("want: %v", wantSorted)
					return
				}
			}
		})
	}
}
