package main

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: 3,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			want: 2,
		},
		{
			name: "Nil tree",
			root: nil,
			want: 0,
		},
		{
			name: "Single node",
			root: &TreeNode{Val: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepth(tt.root)
			if got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
