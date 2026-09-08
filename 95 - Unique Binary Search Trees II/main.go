package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var generate func(left, right int) []*TreeNode
	generate = func(left, right int) []*TreeNode {
		if left > right {
			return []*TreeNode{nil}
		}

		var res []*TreeNode
		for val := left; val <= right; val++ {
			leftTrees := generate(left, val-1)
			rightTrees := generate(val+1, right)

			for _, leftTree := range leftTrees {
				for _, rightTree := range rightTrees {
					res = append(res, &TreeNode{val, leftTree, rightTree})
				}
			}
		}
		return res
	}
	return generate(1, n)
}
