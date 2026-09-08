package main

import (
	"testing"
)

func inorderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	var res []int
	res = append(res, inorderTraversal(node.Left)...)
	res = append(res, node.Val)
	res = append(res, inorderTraversal(node.Right)...)
	return res
}

func preorderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	var res []int
	res = append(res, node.Val)
	res = append(res, preorderTraversal(node.Left)...)
	res = append(res, preorderTraversal(node.Right)...)
	return res
}

func serializeTree(node *TreeNode) string {
	if node == nil {
		return "nil"
	}
	return "[" + serializeTree(node.Left) + "," + string(rune(node.Val+'0')) + "," + serializeTree(node.Right) + "]"
}

func countTrees(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + countTrees(node.Left) + countTrees(node.Right)
}

func isBST(node *TreeNode, minVal, maxVal int) bool {
	if node == nil {
		return true
	}
	if node.Val <= minVal || node.Val >= maxVal {
		return false
	}
	return isBST(node.Left, minVal, node.Val) && isBST(node.Right, node.Val, maxVal)
}

func TestGenerateTrees(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n=1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n=2",
			n:        2,
			expected: 2,
		},
		{
			name:     "n=3",
			n:        3,
			expected: 5,
		},
		{
			name:     "n=4",
			n:        4,
			expected: 14,
		},
		{
			name:     "n=5",
			n:        5,
			expected: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trees := generateTrees(tt.n)
			if len(trees) != tt.expected {
				t.Errorf("expected %d trees, got %d", tt.expected, len(trees))
			}

			inorders := make(map[string]bool)
			for _, tree := range trees {
				if !isBST(tree, 0, tt.n+1) {
					t.Errorf("generated tree is not a valid BST")
				}

				inorder := inorderTraversal(tree)
				if len(inorder) != tt.n {
					t.Errorf("expected inorder length %d, got %d", tt.n, len(inorder))
				}

				key := serializeTree(tree)
				if inorders[key] {
					t.Errorf("duplicate tree found: %s", key)
				}
				inorders[key] = true
			}
		})
	}
}

func TestGenerateTreesStructuralUniqueness(t *testing.T) {
	t.Run("n=1_single_node", func(t *testing.T) {
		trees := generateTrees(1)
		if len(trees) != 1 {
			t.Errorf("expected 1 tree for n=1, got %d", len(trees))
		}
		if trees[0].Val != 1 || trees[0].Left != nil || trees[0].Right != nil {
			t.Errorf("expected single node with value 1")
		}
	})

	t.Run("n=2_all_unique", func(t *testing.T) {
		trees := generateTrees(2)
		if len(trees) != 2 {
			t.Errorf("expected 2 trees for n=2, got %d", len(trees))
		}

		for _, tree := range trees {
			inorder := inorderTraversal(tree)
			if len(inorder) != 2 || (inorder[0] != 1 || inorder[1] != 2) {
				t.Errorf("expected inorder [1, 2], got %v", inorder)
			}
		}
	})
}
