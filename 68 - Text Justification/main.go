package main

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	var current []string
	var currentLen int

	for i := 0; i < len(words); i++ {
		word := words[i]
		if currentLen+len(word)+len(current) > maxWidth {
			result = append(result, justify(current, maxWidth, false))
			current = nil
			currentLen = 0
		}
		current = append(current, word)
		currentLen += len(word)
	}

	if len(current) > 0 {
		result = append(result, justify(current, maxWidth, true))
	}

	return result
}

func justify(words []string, maxWidth int, isLast bool) string {
	if isLast {
		line := ""
		for i, word := range words {
			if i > 0 {
				line += " "
			}
			line += word
		}
		for len(line) < maxWidth {
			line += " "
		}
		return line
	}

	totalLen := 0
	for _, word := range words {
		totalLen += len(word)
	}

	totalSpaces := maxWidth - totalLen
	gaps := len(words) - 1

	if gaps == 0 {
		line := words[0]
		for len(line) < maxWidth {
			line += " "
		}
		return line
	}

	spacePerGap := totalSpaces / gaps
	extraSpaces := totalSpaces % gaps

	line := ""
	for i, word := range words {
		line += word
		if i < gaps {
			spaces := spacePerGap
			if i < extraSpaces {
				spaces++
			}
			for j := 0; j < spaces; j++ {
				line += " "
			}
		}
	}

	return line
}
