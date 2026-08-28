package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev2, prev1 := 1, 2
	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}
	return prev1
}
