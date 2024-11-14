package main

func intToRoman(num int) string {
	// Define os valores e símbolos romanos em pares, do maior para o menor
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// String de resultado para armazenar o número romano
	var result string

	// Itera pelos valores e símbolos
	for i := 0; i < len(values); i++ {
		// Enquanto o número ainda é maior ou igual ao valor atual
		for num >= values[i] {
			num -= values[i]     // Subtrai o valor
			result += symbols[i] // Adiciona o símbolo correspondente ao resultado
		}
	}

	return result
}
