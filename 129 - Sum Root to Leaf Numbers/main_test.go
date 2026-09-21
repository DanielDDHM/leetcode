package main

import "testing"

func arrayToTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		if arr[i] != nil {
			node.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(arr) && arr[i] != nil {
			node.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "example1",
			root:     arrayToTree([]interface{}{1, 2, 3}),
			expected: 25,
		},
		{
			name:     "example2",
			root:     arrayToTree([]interface{}{4, 9, 0, 5, 1}),
			expected: 1026,
		},
		{
			name:     "single_node",
			root:     arrayToTree([]interface{}{5}),
			expected: 5,
		},
		{
			name:     "single_zero",
			root:     arrayToTree([]interface{}{0}),
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumNumbers(tt.root)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
