package main

func combine(n int, k int) [][]int {
	var result [][]int
	var current []int

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(current) == k {
			combo := make([]int, k)
			copy(combo, current)
			result = append(result, combo)
			return
		}

		for i := start; i <= n; i++ {
			current = append(current, i)
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(1)
	return result
}
