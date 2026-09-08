package main

import (
	"slices"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			expected: []int{1, 3, 2},
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{
					Val:  3,
					Left: nil,
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val:   9,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			expected: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			name:     "Example 3: Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name: "Example 4: Single node",
			root: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			expected: []int{1},
		},
		{
			name: "Left-skewed tree",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Right-skewed tree",
			root: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inorderTraversal(tt.root)
			if !slices.Equal(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
