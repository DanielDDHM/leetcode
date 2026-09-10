package main

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "example1",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			want: true,
		},
		{
			name: "example2",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			want: false,
		},
		{
			name: "single_node",
			root: &TreeNode{Val: 1},
			want: true,
		},
		{
			name: "duplicate_values",
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 5},
			},
			want: false,
		},
		{
			name: "valid_bst_larger",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{
					Val:   15,
					Left:  &TreeNode{Val: 12},
					Right: &TreeNode{Val: 20},
				},
			},
			want: true,
		},
		{
			name: "invalid_right_subtree",
			root: &TreeNode{
				Val: 10,
				Right: &TreeNode{
					Val:  15,
					Left: &TreeNode{Val: 6},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidBST(tt.root)
			if got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
