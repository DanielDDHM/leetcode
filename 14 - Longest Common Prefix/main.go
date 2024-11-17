package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Assume que o prefixo inicial é a primeira string
	prefix := strs[0]

	// Itera pelas outras strings
	for i := 1; i < len(strs); i++ {
		// Reduz o prefixo enquanto ele não for um prefixo da string atual
		for len(prefix) > 0 && !startsWith(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1] // Remove o último caractere do prefixo
		}

		// Se o prefixo se tornar vazio, não há prefixo comum
		if len(prefix) == 0 {
			return ""
		}
	}

	return prefix
}

// Função auxiliar para verificar se uma string começa com um prefixo
func startsWith(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return s[:len(prefix)] == prefix
}
