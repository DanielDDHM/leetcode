package main

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	water := 0
	for i := 0; i < n; i++ {
		waterLevel := min(leftMax[i], rightMax[i])
		water += waterLevel - height[i]
	}

	return water
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
