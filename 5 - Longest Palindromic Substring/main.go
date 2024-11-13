package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0

	// Função para expandir em torno de um centro e retornar o tamanho do palíndromo
	expandAroundCenter := func(s string, left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // Expande para palíndromos de comprimento ímpar
		len2 := expandAroundCenter(s, i, i+1) // Expande para palíndromos de comprimento par
		maxLen := len1
		if len2 > len1 {
			maxLen = len2
		}

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func main() {
	fmt.Println(longestPalindrome("babad")) // Output: bab
	fmt.Println(longestPalindrome("cbbd"))  // Output: bb
}
