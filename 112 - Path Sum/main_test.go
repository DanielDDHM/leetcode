package main

import "testing"

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		want      bool
	}{
		{
			name: "Example 1: path exists",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 22,
			want:      true,
		},
		{
			name: "Example 2: no path exists",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			targetSum: 5,
			want:      false,
		},
		{
			name:      "Example 3: empty tree",
			root:      nil,
			targetSum: 0,
			want:      false,
		},
		{
			name: "Single node equals target",
			root: &TreeNode{
				Val: 5,
			},
			targetSum: 5,
			want:      true,
		},
		{
			name: "Single node not equal target",
			root: &TreeNode{
				Val: 5,
			},
			targetSum: 10,
			want:      false,
		},
		{
			name: "Negative values in path",
			root: &TreeNode{
				Val: -2,
				Right: &TreeNode{
					Val: -3,
				},
			},
			targetSum: -5,
			want:      true,
		},
		{
			name: "Path with zero",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -1,
				},
			},
			targetSum: 0,
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasPathSum(tt.root, tt.targetSum)
			if got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
