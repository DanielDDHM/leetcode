package main

func wordBreak(s string, wordDict []string) []string {
	memo := make(map[int][]string)
	return backtrack(s, 0, wordDict, memo)
}

func backtrack(s string, start int, wordDict []string, memo map[int][]string) []string {
	if cached, ok := memo[start]; ok {
		return cached
	}

	var result []string

	if start == len(s) {
		result = append(result, "")
		memo[start] = result
		return result
	}

	for _, word := range wordDict {
		if len(word) <= len(s)-start && s[start:start+len(word)] == word {
			subResults := backtrack(s, start+len(word), wordDict, memo)
			for _, subResult := range subResults {
				if subResult == "" {
					result = append(result, word)
				} else {
					result = append(result, word+" "+subResult)
				}
			}
		}
	}

	memo[start] = result
	return result
}
