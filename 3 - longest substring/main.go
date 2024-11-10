package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int) // mapa para armazenar o índice de cada caractere
	maxLength := 0
	left := 0 // início da janela deslizante

	for right := 0; right < len(s); right++ {
		if index, found := charIndex[s[right]]; found && index >= left {
			// Se o caractere já existe na janela, mova a borda esquerda para o índice após o último caractere duplicado
			left = index + 1
		}

		// Atualiza a posição do caractere atual
		charIndex[s[right]] = right

		// Calcula o comprimento atual da substring e atualiza o comprimento máximo, se necessário
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
}
