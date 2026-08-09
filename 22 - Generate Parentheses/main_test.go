package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	cases := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			name:     "three pairs",
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "one pair",
			n:        1,
			expected: []string{"()"},
		},
		{
			name:     "two pairs",
			n:        2,
			expected: []string{"(())", "()()"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := generateParenthesis(c.n)

			gotSorted := append([]string(nil), got...)
			expectedSorted := append([]string(nil), c.expected...)
			sort.Strings(gotSorted)
			sort.Strings(expectedSorted)

			if !reflect.DeepEqual(gotSorted, expectedSorted) {
				t.Errorf("generateParenthesis(%d) = %v, expected %v", c.n, got, c.expected)
			}
		})
	}
}

func TestGenerateParenthesisCombinationCount(t *testing.T) {
	cases := []struct {
		name  string
		n     int
		count int
	}{
		{name: "smallest allowed n", n: 1, count: 1},
		{name: "four pairs", n: 4, count: 14},
		{name: "largest allowed n", n: 8, count: 1430},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := generateParenthesis(c.n)

			if len(got) != c.count {
				t.Errorf("generateParenthesis(%d) returned %d combinations, expected %d", c.n, len(got), c.count)
			}

			seen := make(map[string]bool, len(got))
			for _, combination := range got {
				if seen[combination] {
					t.Errorf("generateParenthesis(%d) returned duplicate combination %q", c.n, combination)
				}
				seen[combination] = true

				if len(combination) != 2*c.n {
					t.Errorf("generateParenthesis(%d) returned %q with length %d, expected %d", c.n, combination, len(combination), 2*c.n)
				}

				if !isWellFormed(combination) {
					t.Errorf("generateParenthesis(%d) returned malformed combination %q", c.n, combination)
				}
			}
		})
	}
}

func isWellFormed(combination string) bool {
	balance := 0

	for _, character := range combination {
		if character == '(' {
			balance++
		} else {
			balance--
		}

		if balance < 0 {
			return false
		}
	}

	return balance == 0
}
