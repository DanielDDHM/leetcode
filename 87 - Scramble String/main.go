package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	n := len(s1)
	memo := make(map[string]bool)

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == 1 {
			return s1[i] == s2[j]
		}

		key := fmt.Sprintf("%d,%d,%d", i, j, k)
		if v, exists := memo[key]; exists {
			return v
		}

		if s1[i:i+k] == s2[j:j+k] {
			memo[key] = true
			return true
		}

		for h := 1; h < k; h++ {
			if (dfs(i, j, h) && dfs(i+h, j+h, k-h)) ||
				(dfs(i, j+k-h, h) && dfs(i+h, j, k-h)) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false
		return false
	}

	return dfs(0, 0, n)
}
