package main

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "Example 1: balanced tree",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
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
			want: true,
		},
		{
			name: "Example 2: unbalanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			want: false,
		},
		{
			name: "Example 3: empty tree",
			root: nil,
			want: true,
		},
		{
			name: "Edge case: single node",
			root: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			want: true,
		},
		{
			name: "Edge case: completely unbalanced left tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBalanced(tt.root)
			if got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
