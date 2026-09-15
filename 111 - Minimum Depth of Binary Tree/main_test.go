package main

import "testing"

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "nil tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "single node",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name: "example 1",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 2,
		},
		{
			name: "example 2",
			root: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val:   5,
							Right: &TreeNode{Val: 6},
						},
					},
				},
			},
			expected: 5,
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
			expected: 3,
		},
		{
			name: "balanced tree depth 2",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minDepth(tt.root)
			if result != tt.expected {
				t.Errorf("minDepth() = %d, want %d", result, tt.expected)
			}
		})
	}
}
