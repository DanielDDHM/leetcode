package main

import (
	"reflect"
	"testing"
)

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}
	traverse(root)
	return result
}

func TestRecoverTree(t *testing.T) {
	tests := []struct {
		name     string
		input    *TreeNode
		expected []int
	}{
		{
			name: "example 1",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 2},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "example 2",
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "two nodes swapped at ends",
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "two adjacent nodes swapped",
			input: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.input)
			result := inorderTraversal(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
