package main

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	current := make([]byte, 0, 2*n)

	var backtrack func(opened, closed int)
	backtrack = func(opened, closed int) {
		if len(current) == 2*n {
			result = append(result, string(current))
			return
		}

		if opened < n {
			current = append(current, '(')
			backtrack(opened+1, closed)
			current = current[:len(current)-1]
		}

		if closed < opened {
			current = append(current, ')')
			backtrack(opened, closed+1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0, 0)

	return result
}
