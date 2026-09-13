package main

import (
	"testing"
)

func treeToSlice(node *TreeNode) []interface{} {
	if node == nil {
		return []interface{}{}
	}
	result := make([]interface{}, 0)
	queue := []*TreeNode{node}
	for len(queue) > 0 {
		newQueue := make([]*TreeNode, 0)
		hasNode := false
		for _, n := range queue {
			if n != nil {
				result = append(result, n.Val)
				newQueue = append(newQueue, n.Left)
				newQueue = append(newQueue, n.Right)
				hasNode = true
			} else {
				result = append(result, nil)
				newQueue = append(newQueue, nil)
				newQueue = append(newQueue, nil)
			}
		}
		if !hasNode {
			break
		}
		queue = newQueue
	}

	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}
	return result
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name      string
		inorder   []int
		postorder []int
		expected  []interface{}
	}{
		{
			name:      "Example 1",
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			expected:  []interface{}{3, 9, 20, nil, nil, 15, 7},
		},
		{
			name:      "Example 2",
			inorder:   []int{-1},
			postorder: []int{-1},
			expected:  []interface{}{-1},
		},
		{
			name:      "Left skewed tree",
			inorder:   []int{3, 2, 1},
			postorder: []int{3, 2, 1},
			expected:  []interface{}{1, 2, nil, 3},
		},
		{
			name:      "Right skewed tree",
			inorder:   []int{1, 2, 3},
			postorder: []int{3, 2, 1},
			expected:  []interface{}{1, nil, 2, nil, nil, nil, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildTree(tt.inorder, tt.postorder)
			resultSlice := treeToSlice(result)

			if len(resultSlice) != len(tt.expected) {
				t.Fatalf("length mismatch: got %d, expected %d", len(resultSlice), len(tt.expected))
			}

			for i, v := range resultSlice {
				if v != tt.expected[i] {
					t.Errorf("at index %d: got %v, expected %v", i, v, tt.expected[i])
				}
			}
		})
	}
}
