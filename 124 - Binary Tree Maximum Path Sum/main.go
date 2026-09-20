package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := root.Val

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := max(0, dfs(node.Left))
		rightSum := max(0, dfs(node.Right))

		maxSum = max(maxSum, node.Val+leftSum+rightSum)

		return node.Val + max(leftSum, rightSum)
	}

	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
