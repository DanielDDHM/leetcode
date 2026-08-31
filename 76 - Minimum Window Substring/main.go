package main

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	charCount := make(map[rune]int)
	for _, ch := range t {
		charCount[ch]++
	}

	required := len(charCount)
	formed := 0

	windowCounts := make(map[rune]int)
	left := 0
	minLen := len(s) + 1
	minLeft := 0

	for right := 0; right < len(s); right++ {
		ch := rune(s[right])
		windowCounts[ch]++

		if count, exists := charCount[ch]; exists && windowCounts[ch] == count {
			formed++
		}

		for formed == required && left <= right {
			if right-left+1 < minLen {
				minLen = right - left + 1
				minLeft = left
			}

			leftCh := rune(s[left])
			if count, exists := charCount[leftCh]; exists && windowCounts[leftCh] == count {
				formed--
			}
			windowCounts[leftCh]--
			left++
		}
	}

	if minLen > len(s) {
		return ""
	}

	return s[minLeft : minLeft+minLen]
}
