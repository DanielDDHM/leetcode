package main

import (
	"slices"
	"testing"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		expected  [][]int
	}{
		{
			name: "example 1",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{Val: 13},
					Right: &TreeNode{
						Val:   4,
						Right: &TreeNode{Val: 5},
					},
				},
			},
			targetSum: 22,
			expected:  [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			name:      "nil root",
			root:      nil,
			targetSum: 0,
			expected:  [][]int{},
		},
		{
			name:      "single node matching",
			root:      &TreeNode{Val: 1},
			targetSum: 1,
			expected:  [][]int{{1}},
		},
		{
			name:      "single node not matching",
			root:      &TreeNode{Val: 1},
			targetSum: 2,
			expected:  [][]int{},
		},
		{
			name: "negative values",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   -2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 1},
			},
			targetSum: 2,
			expected:  [][]int{{1, -2, 3}, {1, 1}},
		},
		{
			name: "left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			targetSum: 6,
			expected:  [][]int{{1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pathSum(tt.root, tt.targetSum)
			if len(result) != len(tt.expected) {
				t.Errorf("got %d paths, expected %d", len(result), len(tt.expected))
				return
			}

			for _, expPath := range tt.expected {
				found := false
				for _, resPath := range result {
					if slices.Equal(expPath, resPath) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected path %v not found in result %v", expPath, result)
				}
			}
		})
	}
}
