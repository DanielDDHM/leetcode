package main

import (
	"testing"
)

func treeToList(node *TreeNode) []interface{} {
	if node == nil {
		return []interface{}{nil}
	}
	result := []interface{}{}
	queue := []*TreeNode{node}
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

func isHeightBalanced(node *TreeNode) bool {
	if node == nil {
		return true
	}
	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)
	if leftHeight < 0 || rightHeight < 0 {
		return false
	}
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		return false
	}
	return true
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)
	if leftHeight < 0 || rightHeight < 0 {
		return -1
	}
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func isBST(node *TreeNode, min, max int, first bool) bool {
	if node == nil {
		return true
	}
	if !first && (node.Val <= min || node.Val >= max) {
		return false
	}
	return isBST(node.Left, min, node.Val, false) && isBST(node.Right, node.Val, max, false)
}

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "example1_negative_and_positive",
			nums: []int{-10, -3, 0, 5, 9},
		},
		{
			name: "example2_two_elements",
			nums: []int{1, 3},
		},
		{
			name: "single_element",
			nums: []int{0},
		},
		{
			name: "two_elements_close",
			nums: []int{1, 2},
		},
		{
			name: "three_elements",
			nums: []int{1, 2, 3},
		},
		{
			name: "all_negative",
			nums: []int{-10, -5, -3, -1},
		},
		{
			name: "all_positive",
			nums: []int{1, 3, 5, 7},
		},
		{
			name: "large_range",
			nums: []int{-10000, -9999, 0, 9999, 10000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedArrayToBST(tt.nums)

			if !isBST(root, -100001, 100001, true) {
				t.Error("Result is not a valid BST")
			}

			if !isHeightBalanced(root) {
				t.Error("Result is not height-balanced")
			}
		})
	}
}
