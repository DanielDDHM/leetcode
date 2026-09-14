package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(int, int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		return &TreeNode{
			Val:   nums[mid],
			Left:  helper(left, mid-1),
			Right: helper(mid+1, right),
		}
	}
	return helper(0, len(nums)-1)
}
