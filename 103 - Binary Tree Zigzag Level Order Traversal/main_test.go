package main

import (
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Example 1",
			root:     buildTree([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name:     "Example 2",
			root:     buildTree([]interface{}{1}),
			expected: [][]int{{1}},
		},
		{
			name:     "Example 3",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "Single left child",
			root:     buildTree([]interface{}{1, 2}),
			expected: [][]int{{1}, {2}},
		},
		{
			name:     "Complete binary tree",
			root:     buildTree([]interface{}{1, 2, 3, 4, 5, 6, 7}),
			expected: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := zigzagLevelOrder(tt.root)
			if !equalSlices(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func buildTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func equalSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
