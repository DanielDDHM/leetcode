package main

import "testing"

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "example 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: 6,
		},
		{
			name: "example 2",
			root: &TreeNode{
				Val: -10,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: 42,
		},
		{
			name: "single node",
			root: &TreeNode{
				Val: 5,
			},
			want: 5,
		},
		{
			name: "all negative",
			root: &TreeNode{
				Val: -3,
				Left: &TreeNode{
					Val: -2,
				},
				Right: &TreeNode{
					Val: -1,
				},
			},
			want: -1,
		},
		{
			name: "negative root with positive children",
			root: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
