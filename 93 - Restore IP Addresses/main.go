package main

func restoreIpAddresses(s string) (ans []string) {
	n := len(s)
	if n < 4 || n > 12 {
		return
	}

	var path []string

	var dfs func(idx int)
	dfs = func(idx int) {
		if len(path) == 4 {
			if idx == n {
				ans = append(ans, join(path))
			}
			return
		}

		if idx >= n {
			return
		}

		remainingDigits := n - idx
		remainingSegments := 4 - len(path)
		if remainingDigits < remainingSegments || remainingDigits > remainingSegments*3 {
			return
		}

		for segLen := 1; segLen <= 3 && idx+segLen <= n; segLen++ {
			segment := s[idx : idx+segLen]

			if !isValid(segment) {
				continue
			}

			path = append(path, segment)
			dfs(idx + segLen)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return
}

func isValid(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}

	if len(s) > 1 && s[0] == '0' {
		return false
	}

	num := 0
	for _, c := range s {
		num = num*10 + int(c-'0')
	}

	return num >= 0 && num <= 255
}

func join(segments []string) string {
	result := ""
	for i, seg := range segments {
		if i > 0 {
			result += "."
		}
		result += seg
	}
	return result
}

func main() {}
