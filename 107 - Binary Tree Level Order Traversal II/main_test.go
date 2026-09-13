package main

import (
	"reflect"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name: "example 1: standard tree",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			name:     "empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "single node",
			root:     &TreeNode{Val: 1},
			expected: [][]int{{1}},
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
			expected: [][]int{{3}, {2}, {1}},
		},
		{
			name: "right skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: [][]int{{3}, {2}, {1}},
		},
		{
			name: "complete binary tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: [][]int{{4, 5, 6, 7}, {2, 3}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrderBottom(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("levelOrderBottom() = %v, want %v", result, tt.expected)
			}
		})
	}
}
