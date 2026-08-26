package main

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]

	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = grid[i][j] + min(dp[j], dp[j-1])
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
