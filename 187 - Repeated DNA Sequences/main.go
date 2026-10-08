package main

func findRepeatedDnaSequences(s string) []string {
	const windowSize = 10
	seen := make(map[string]bool)
	result := make(map[string]bool)

	for i := 0; i+windowSize <= len(s); i++ {
		window := s[i : i+windowSize]
		if seen[window] {
			result[window] = true
		} else {
			seen[window] = true
		}
	}

	ans := make([]string, 0, len(result))
	for seq := range result {
		ans = append(ans, seq)
	}
	return ans
}
