package main

import (
	"fmt"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}

	if !wordSet[endWord] {
		return [][]string{}
	}

	neighbors := make(map[string][]string)
	distance := make(map[string]int)

	bfs(beginWord, endWord, wordSet, neighbors, distance)

	result := [][]string{}
	path := []string{beginWord}
	dfs(beginWord, endWord, neighbors, distance, path, &result)

	return result
}

func bfs(beginWord, endWord string, wordSet map[string]bool, neighbors map[string][]string, distance map[string]int) {
	for word := range wordSet {
		distance[word] = -1
	}
	distance[beginWord] = 0

	queue := []string{beginWord}
	for len(queue) > 0 {
		nextQueue := []string{}
		for _, word := range queue {
			if distance[word] == -1 {
				continue
			}
			neighbors[word] = []string{}
			for next := range getNextWords(word, wordSet) {
				neighbors[word] = append(neighbors[word], next)
				if distance[next] == -1 {
					distance[next] = distance[word] + 1
					nextQueue = append(nextQueue, next)
				}
			}
		}
		queue = nextQueue
	}
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

func dfs(word, endWord string, neighbors map[string][]string, distance map[string]int, path []string, result *[][]string) {
	if word == endWord {
		*result = append(*result, append([]string{}, path...))
		return
	}

	for _, neighbor := range neighbors[word] {
		if distance[neighbor] == distance[word]+1 {
			dfs(neighbor, endWord, neighbors, distance, append(path, neighbor), result)
		}
	}
}

func main() {
	fmt.Println("LeetCode 126: Word Ladder II")
}
