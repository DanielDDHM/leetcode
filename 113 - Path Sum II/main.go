package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int
	dfs(root, targetSum, 0, &path, &result)
	return result
}

func dfs(node *TreeNode, target int, currentSum int, path *[]int, result *[][]int) {
	if node == nil {
		return
	}

	currentSum += node.Val
	*path = append(*path, node.Val)

	if node.Left == nil && node.Right == nil {
		if currentSum == target {
			*result = append(*result, slices.Clone(*path))
		}
	} else {
		dfs(node.Left, target, currentSum, path, result)
		dfs(node.Right, target, currentSum, path, result)
	}

	*path = (*path)[:len(*path)-1]
}
