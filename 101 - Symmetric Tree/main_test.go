package main

import "testing"

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name: "Example 1 - symmetric tree with 7 nodes",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 3},
				},
			},
			expected: true,
		},
		{
			name: "Example 2 - asymmetric tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: false,
		},
		{
			name:     "Single node tree",
			root:     &TreeNode{Val: 1},
			expected: true,
		},
		{
			name: "Symmetric with different values",
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -1,
				},
				Right: &TreeNode{
					Val: -1,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSymmetric(tt.root)
			if result != tt.expected {
				t.Errorf("isSymmetric() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
