package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		key := sortString(str)
		m[key] = append(m[key], str)
	}

	result := make([][]string, 0, len(m))
	for _, group := range m {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
