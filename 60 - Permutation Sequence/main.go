package main

import "strconv"

func getPermutation(n int, k int) string {
	factorials := make([]int, n)
	factorials[0] = 1
	for i := 1; i < n; i++ {
		factorials[i] = factorials[i-1] * i
	}

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i + 1
	}

	k--
	result := ""
	for i := 0; i < n; i++ {
		fact := factorials[n-1-i]
		idx := k / fact
		result += strconv.Itoa(numbers[idx])
		numbers = append(numbers[:idx], numbers[idx+1:]...)
		k %= fact
	}

	return result
}
