package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	n := len(s)
	prev2, prev1 := 1, 1

	for i := 1; i < n; i++ {
		current := 0

		if s[i] != '0' {
			current = prev1
		}

		twoDigit := (s[i-1]-'0')*10 + (s[i] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			current += prev2
		}

		if current == 0 {
			return 0
		}

		prev2, prev1 = prev1, current
	}

	return prev1
}
