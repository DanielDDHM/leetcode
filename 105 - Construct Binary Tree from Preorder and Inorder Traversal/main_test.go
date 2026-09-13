package main

import (
	"testing"
)

func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	var result []interface{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}

	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		expected []interface{}
	}{
		{
			name:     "Example 1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expected: []interface{}{3, 9, 20, nil, nil, 15, 7},
		},
		{
			name:     "Example 2",
			preorder: []int{-1},
			inorder:  []int{-1},
			expected: []interface{}{-1},
		},
		{
			name:     "Left skewed tree",
			preorder: []int{3, 2, 1},
			inorder:  []int{1, 2, 3},
			expected: []interface{}{3, 2, nil, 1},
		},
		{
			name:     "Right skewed tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{1, 2, 3},
			expected: []interface{}{1, nil, 2, nil, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.preorder, tt.inorder)
			result := treeToSlice(root)

			if len(result) != len(tt.expected) {
				t.Errorf("Length mismatch. Expected %d, got %d", len(tt.expected), len(result))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Value mismatch at index %d. Expected %v, got %v", i, tt.expected[i], result[i])
				}
			}
		})
	}
}
