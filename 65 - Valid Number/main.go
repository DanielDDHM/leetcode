package main

func isNumber(s string) bool {
	i := 0
	n := len(s)

	for i < n && s[i] == ' ' {
		i++
	}

	if i >= n {
		return false
	}

	if s[i] == '+' || s[i] == '-' {
		i++
	}

	hasDigitBeforeDot := false
	for i < n && s[i] >= '0' && s[i] <= '9' {
		hasDigitBeforeDot = true
		i++
	}

	hasDigitAfterDot := false
	if i < n && s[i] == '.' {
		i++
		for i < n && s[i] >= '0' && s[i] <= '9' {
			hasDigitAfterDot = true
			i++
		}
	}

	if !hasDigitBeforeDot && !hasDigitAfterDot {
		return false
	}

	if i < n && (s[i] == 'e' || s[i] == 'E') {
		i++

		if i >= n {
			return false
		}

		if s[i] == '+' || s[i] == '-' {
			i++
		}

		if i >= n {
			return false
		}

		hasDigitInExp := false
		for i < n && s[i] >= '0' && s[i] <= '9' {
			hasDigitInExp = true
			i++
		}

		if !hasDigitInExp {
			return false
		}
	}

	for i < n && s[i] == ' ' {
		i++
	}

	return i == n
}
