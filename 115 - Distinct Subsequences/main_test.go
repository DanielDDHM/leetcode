package main

import "testing"

func TestNumDistinct(t *testing.T) {
	cases := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{"Example1", "rabbbit", "rabbit", 3},
		{"Example2", "babgbag", "bag", 5},
		{"SingleChar", "a", "a", 1},
		{"NoMatch", "a", "b", 0},
		{"MultipleChars", "aaa", "a", 3},
		{"TGreaterThanS", "a", "aa", 0},
		{"LongSequence", "aabaab", "aab", 7},
		{"NoCommon", "xyz", "abc", 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := numDistinct(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("numDistinct(%q, %q) = %d, want %d", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
