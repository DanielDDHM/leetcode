package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// Criar slices para cada linha
	rows := make([]string, numRows)
	curRow := 0
	goingDown := false

	// Distribuir os caracteres em cada linha de acordo com o padrão zigzag
	for _, char := range s {
		rows[curRow] += string(char)
		// Inverter a direção quando chegar na primeira ou última linha
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	// Concatenar todas as linhas para formar o resultado
	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}

func main() {
	// Casos de teste
	testCases := []struct {
		s       string
		numRows int
		expect  string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
		{"AB", 1, "AB"},
	}

	// Executa cada caso de teste e exibe o resultado
	for _, tc := range testCases {
		result := convert(tc.s, tc.numRows)
		fmt.Printf("convert(%q, %d) = %q; Expected: %q\n", tc.s, tc.numRows, result, tc.expect)
	}
}
