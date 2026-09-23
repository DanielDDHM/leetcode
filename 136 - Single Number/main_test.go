package main

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			name:   "example1",
			nums:   []int{2, 2, 1},
			expect: 1,
		},
		{
			name:   "example2",
			nums:   []int{4, 1, 2, 1, 2},
			expect: 4,
		},
		{
			name:   "example3",
			nums:   []int{1},
			expect: 1,
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -1, 0},
			expect: 0,
		},
		{
			name:   "large array with single at start",
			nums:   []int{5, 2, 2, 3, 3},
			expect: 5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := singleNumber(tc.nums)
			if result != tc.expect {
				t.Errorf("got %d, expect %d", result, tc.expect)
			}
		})
	}
}
