package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	// Remover espaços em branco
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	// Variáveis para sinal e índice
	sign := 1
	index := 0

	// Verificar o sinal
	if s[0] == '-' {
		sign = -1
		index++
	} else if s[0] == '+' {
		index++
	}

	// Converter dígitos em número inteiro
	result := 0
	for index < len(s) && s[index] >= '0' && s[index] <= '9' {
		// Obter o dígito atual
		digit := int(s[index] - '0')

		// Verificar overflow antes de multiplicar
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		// Atualizar o resultado
		result = result*10 + digit
		index++
	}

	return result * sign
}

func main() {
	// Testes básicos
	fmt.Println(myAtoi("42"))              // Esperado: 42
	fmt.Println(myAtoi("   -42"))          // Esperado: -42
	fmt.Println(myAtoi("4193 with words")) // Esperado: 4193
	fmt.Println(myAtoi("words and 987"))   // Esperado: 0
	fmt.Println(myAtoi("-91283472332"))    // Esperado: -2147483648 (limite inferior de 32 bits)

	// Testes adicionais
	fmt.Println(myAtoi("+123"))        // Esperado: 123
	fmt.Println(myAtoi("0000012345"))  // Esperado: 12345
	fmt.Println(myAtoi("2147483647"))  // Esperado: 2147483647 (limite superior de 32 bits)
	fmt.Println(myAtoi("2147483648"))  // Esperado: 2147483647 (limite superior de 32 bits)
	fmt.Println(myAtoi("-2147483648")) // Esperado: -2147483648
	fmt.Println(myAtoi("-2147483649")) // Esperado: -2147483648 (limite inferior de 32 bits)
}
