package main

func romanToInt(s string) int {
	// Mapa para os valores dos numerais romanos
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		// Pega o valor atual
		current := romanMap[s[i]]

		// Verifica se precisa subtrair (casos como IV, IX, etc.)
		if i < n-1 && current < romanMap[s[i+1]] {
			total -= current
		} else {
			total += current
		}
	}

	return total
}
