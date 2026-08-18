package main

import "testing"

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "three elements",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name: "two elements",
			nums: []int{0, 1},
			want: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name: "single element",
			nums: []int{1},
			want: [][]int{
				{1},
			},
		},
		{
			name: "two elements with negative",
			nums: []int{-1, 2},
			want: [][]int{
				{-1, 2},
				{2, -1},
			},
		},
		{
			name: "four elements",
			nums: []int{1, 2, 3, 4},
			want: generateExpectedPermutations([]int{1, 2, 3, 4}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			if !permutationsEqual(got, tt.want) {
				t.Errorf("permute(%v) got %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func permutationsEqual(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}

	seen := make(map[string]bool)
	for _, perm := range want {
		key := sliceToString(perm)
		seen[key] = true
	}

	for _, perm := range got {
		key := sliceToString(perm)
		if !seen[key] {
			return false
		}
	}

	return true
}

func sliceToString(nums []int) string {
	var result string
	for i, v := range nums {
		if i > 0 {
			result += ","
		}
		result += string(rune(v))
	}
	return result
}

func generateExpectedPermutations(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, &result, used)
	return result
}
