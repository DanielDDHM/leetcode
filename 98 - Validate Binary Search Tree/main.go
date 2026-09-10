package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validateHelper(root, nil, nil)
}

func validateHelper(node *TreeNode, minVal, maxVal *int) bool {
	if node == nil {
		return true
	}

	if minVal != nil && node.Val <= *minVal {
		return false
	}
	if maxVal != nil && node.Val >= *maxVal {
		return false
	}

	return validateHelper(node.Left, minVal, &node.Val) && validateHelper(node.Right, &node.Val, maxVal)
}
