package main

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 9}},
				},
			},
			want: []int{4, 6, 7, 5, 2, 9, 8, 3, 1},
		},
		{
			name: "Example 3 - Empty tree",
			root: nil,
			want: []int{},
		},
		{
			name: "Example 4 - Single node",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "Only right children",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "Only left children",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
			},
			want: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := postorderTraversal(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
