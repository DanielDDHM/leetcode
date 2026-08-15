package main

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int

	var backtrack func(int, int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			combination := make([]int, len(current))
			copy(combination, current)
			result = append(result, combination)
			return
		}

		if remaining < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			current = append(current, candidates[i])
			backtrack(i, remaining-candidates[i])
			current = current[:len(current)-1]
		}
	}

	backtrack(0, target)
	return result
}
