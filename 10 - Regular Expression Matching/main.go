package main

func isMatch(s string, p string) bool {
	// `m` é o comprimento da string `s`, e `n` é o comprimento do padrão `p`
	m, n := len(s), len(p)

	// Criamos uma matriz dp de tamanho (m+1) x (n+1), onde dp[i][j] indica se s[0:i] corresponde a p[0:j]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// A posição dp[0][0] é verdadeira porque uma string vazia corresponde a um padrão vazio
	dp[0][0] = true

	// Preenche as correspondências com padrões que contêm `*`, que podem corresponder a uma string vazia
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2] // `*` pode remover o caractere anterior
		}
	}

	// Iteramos sobre cada caractere de `s` e `p` para preencher a matriz dp
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				// Caso 1: se os caracteres são iguais ou `p[j-1]` é `.`, herda a correspondência de dp[i-1][j-1]
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// Caso 2: `*` pode ignorar o caractere anterior ou "repeti-lo"
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
				// dp[i][j-2]: ignora o par "<char>*"
				// dp[i-1][j]: corresponde ao caractere atual de `s` com `<char>*`
			}
		}
	}

	// O resultado final está em dp[m][n], indicando se `s` corresponde a `p`
	return dp[m][n]
}
