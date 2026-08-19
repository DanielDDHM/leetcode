package main

import (
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name: "Example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name:     "Example 2",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Example 3",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "All different non-anagrams",
			strs:     []string{"ab", "cd", "ef"},
			expected: [][]string{{"ab"}, {"cd"}, {"ef"}},
		},
		{
			name:     "All same strings",
			strs:     []string{"hello", "hello", "hello"},
			expected: [][]string{{"hello", "hello", "hello"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.strs)
			if !anagramGroupsEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func anagramGroupsEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	sortGroups := func(groups [][]string) [][]string {
		copied := make([][]string, len(groups))
		for i, group := range groups {
			copied[i] = make([]string, len(group))
			copy(copied[i], group)
		}

		for _, group := range copied {
			sort.Strings(group)
		}

		sort.Slice(copied, func(i, j int) bool {
			if len(copied[i]) != len(copied[j]) {
				return len(copied[i]) < len(copied[j])
			}
			for k := range copied[i] {
				if copied[i][k] != copied[j][k] {
					return copied[i][k] < copied[j][k]
				}
			}
			return false
		})

		return copied
	}

	return isEqual(sortGroups(a), sortGroups(b))
}

func isEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
