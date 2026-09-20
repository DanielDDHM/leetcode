package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	wordSet[beginWord] = true
	queue := []string{beginWord}
	distance := 1

	for len(queue) > 0 {
		nextQueue := []string{}
		for _, word := range queue {
			if word == endWord {
				return distance
			}

			for next := range getNextWords(word, wordSet) {
				nextQueue = append(nextQueue, next)
				delete(wordSet, next)
			}
		}
		queue = nextQueue
		distance++
	}

	return 0
}

func getNextWords(word string, wordSet map[string]bool) map[string]bool {
	neighbors := make(map[string]bool)
	chars := []byte(word)

	for i := 0; i < len(chars); i++ {
		old := chars[i]
		for c := byte('a'); c <= 'z'; c++ {
			chars[i] = c
			newWord := string(chars)
			if wordSet[newWord] {
				neighbors[newWord] = true
			}
		}
		chars[i] = old
	}

	return neighbors
}
