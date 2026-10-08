package main

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[m-1][n-1] = max(1, 1-dungeon[m-1][n-1])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}

			var minHealthNeeded int
			canGoRight := j+1 < n
			canGoDown := i+1 < m

			if canGoRight && canGoDown {
				minHealthNeeded = min(max(1, dp[i][j+1]-dungeon[i][j]), max(1, dp[i+1][j]-dungeon[i][j]))
			} else if canGoRight {
				minHealthNeeded = max(1, dp[i][j+1]-dungeon[i][j])
			} else if canGoDown {
				minHealthNeeded = max(1, dp[i+1][j]-dungeon[i][j])
			}

			dp[i][j] = minHealthNeeded
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
