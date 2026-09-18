package main

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = triangle[n-1][i]
	}

	for row := n - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			if dp[col] > dp[col+1] {
				dp[col] = dp[col+1]
			}
			dp[col] += triangle[row][col]
		}
	}

	return dp[0]
}
