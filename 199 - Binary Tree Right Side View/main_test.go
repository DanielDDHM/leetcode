package main

import (
	"slices"
	"testing"
)

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

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name   string
		root   *TreeNode
		expect []int
	}{
		{
			name:   "Example 1",
			root:   buildTree([]interface{}{1, 2, 3, nil, 5, nil, 4}),
			expect: []int{1, 3, 4},
		},
		{
			name:   "Example 2",
			root:   buildTree([]interface{}{1, 2, 3, 4, nil, nil, nil, 5}),
			expect: []int{1, 3, 4, 5},
		},
		{
			name:   "Example 3",
			root:   buildTree([]interface{}{1, nil, 3}),
			expect: []int{1, 3},
		},
		{
			name:   "Example 4",
			root:   buildTree([]interface{}{}),
			expect: []int{},
		},
		{
			name:   "Single node",
			root:   buildTree([]interface{}{1}),
			expect: []int{1},
		},
		{
			name:   "Left-only tree",
			root:   buildTree([]interface{}{1, 2, nil, 3}),
			expect: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rightSideView(tt.root)
			if !slices.Equal(result, tt.expect) {
				t.Errorf("got %v, want %v", result, tt.expect)
			}
		})
	}
}
