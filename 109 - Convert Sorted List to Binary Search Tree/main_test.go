package main

import (
	"testing"
)

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	return head
}

func treeToInorder(node *TreeNode) []int {
	var result []int
	var inorder func(*TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inorder(n.Left)
		result = append(result, n.Val)
		inorder(n.Right)
	}
	inorder(node)
	return result
}

func isBalanced(node *TreeNode) bool {
	_, ok := checkBalance(node)
	return ok
}

func checkBalance(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftBalanced := checkBalance(node.Left)
	if !leftBalanced {
		return 0, false
	}

	rightHeight, rightBalanced := checkBalance(node.Right)
	if !rightBalanced {
		return 0, false
	}

	if leftHeight-rightHeight < -1 || leftHeight-rightHeight > 1 {
		return 0, false
	}

	height := 1
	if leftHeight > rightHeight {
		height += leftHeight
	} else {
		height += rightHeight
	}

	return height, true
}

func isBST(node *TreeNode) bool {
	return isBSTHelper(node, nil, nil)
}

func isBSTHelper(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}

	return isBSTHelper(node.Left, min, &node.Val) &&
		isBSTHelper(node.Right, &node.Val, max)
}

func TestSortedListToBST(t *testing.T) {
	tests := []struct {
		name   string
		values []int
	}{
		{
			name:   "Example 1",
			values: []int{-10, -3, 0, 5, 9},
		},
		{
			name:   "Example 2",
			values: []int{},
		},
		{
			name:   "Single element",
			values: []int{1},
		},
		{
			name:   "Two elements",
			values: []int{1, 2},
		},
		{
			name:   "Three elements",
			values: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.values)
			result := sortedListToBST(head)

			inorder := treeToInorder(result)
			if len(inorder) != len(tt.values) {
				t.Errorf("Expected %d nodes, got %d", len(tt.values), len(inorder))
			}

			for i := 0; i < len(inorder); i++ {
				if inorder[i] != tt.values[i] {
					t.Errorf("Inorder traversal mismatch at index %d: expected %d, got %d",
						i, tt.values[i], inorder[i])
				}
			}

			if result != nil && !isBalanced(result) {
				t.Errorf("Tree is not balanced")
			}

			if result != nil && !isBST(result) {
				t.Errorf("Tree is not a valid BST")
			}
		})
	}
}
