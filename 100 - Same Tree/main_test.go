package main

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			name: "Example 1: Both identical trees with same structure",
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want: true,
		},
		{
			name: "Example 2: Different structures",
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			q:    &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			want: false,
		},
		{
			name: "Example 3: Same structure but different values",
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			q:    &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			want: false,
		},
		{
			name: "Edge case: Both trees are nil",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "Edge case: One tree is nil, other is not",
			p:    &TreeNode{Val: 1},
			q:    nil,
			want: false,
		},
		{
			name: "Edge case: Single node trees with same value",
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 5},
			want: true,
		},
		{
			name: "Edge case: Single node trees with different values",
			p:    &TreeNode{Val: 1},
			q:    &TreeNode{Val: 2},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(tt.p, tt.q)
			if got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
