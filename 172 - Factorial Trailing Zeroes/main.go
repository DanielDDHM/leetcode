package main

func trailingZeroes(n int) int {
	count := 0
	divisor := 5
	for divisor <= n {
		count += n / divisor
		divisor *= 5
	}
	return count
}
