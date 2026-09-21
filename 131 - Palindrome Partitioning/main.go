package main

func partition(s string) [][]string {
	n := len(s)

	isPalin := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalin[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || isPalin[i+1][j-1]) {
				isPalin[i][j] = true
			}
		}
	}

	var result [][]string
	var current []string

	var backtrack func(int)
	backtrack = func(start int) {
		if start == n {
			result = append(result, append([]string{}, current...))
			return
		}

		for end := start; end < n; end++ {
			if isPalin[start][end] {
				current = append(current, s[start:end+1])
				backtrack(end + 1)
				current = current[:len(current)-1]
			}
		}
	}

	backtrack(0)
	return result
}
