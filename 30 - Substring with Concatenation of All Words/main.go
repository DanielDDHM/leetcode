package main

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordCount := len(words)
	windowSize := wordLen * wordCount
	result := []int{}

	if len(s) < windowSize {
		return result
	}

	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	for i := 0; i <= len(s)-windowSize; i++ {
		window := s[i : i+windowSize]
		windowFreq := make(map[string]int)

		for j := 0; j < windowSize; j += wordLen {
			word := window[j : j+wordLen]
			windowFreq[word]++
		}

		if mapsEqual(wordFreq, windowFreq) {
			result = append(result, i)
		}
	}

	return result
}

func mapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, val := range a {
		if b[key] != val {
			return false
		}
	}
	return true
}
