package main

func minCut(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}

	isPalin := make([][]bool, n)
	for i := range isPalin {
		isPalin[i] = make([]bool, n)
		isPalin[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			isPalin[i][i+1] = true
		}
	}

	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && isPalin[i+1][j-1] {
				isPalin[i][j] = true
			}
		}
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = i
		for j := 0; j <= i; j++ {
			if isPalin[j][i] {
				if j == 0 {
					dp[i] = 0
				} else {
					dp[i] = min(dp[i], dp[j-1]+1)
				}
			}
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
