package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, ok := checkBalance(root)
	return ok
}

func checkBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHeight, leftBalanced := checkBalance(root.Left)
	if !leftBalanced {
		return 0, false
	}

	rightHeight, rightBalanced := checkBalance(root.Right)
	if !rightBalanced {
		return 0, false
	}

	diff := leftHeight - rightHeight
	if diff < 0 {
		diff = -diff
	}

	if diff > 1 {
		return 0, false
	}

	height := leftHeight
	if rightHeight > leftHeight {
		height = rightHeight
	}

	return height + 1, true
}
