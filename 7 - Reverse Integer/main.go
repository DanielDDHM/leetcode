package main

import "fmt"

func reverse(x int) int {
	var result int
	for x != 0 {
		reminder := x % 10
		x /= 10

		// Verifica se o resultado está fora dos limites de 32 bits para inteiros com sinal
		if result > (1<<31-1)/10 || (result == (1<<31-1)/10 && reminder > 7) {
			return 0
		}
		if result < (-1<<31)/10 || (result == (-1<<31)/10 && reminder < -8) {
			return 0
		}
		result = result*10 + reminder
	}
	return result
}

func main() {
	tests := []int{123, -123, 120, 1534236469, -2147483648}

	for _, test := range tests {
		result := reverse(test)
		fmt.Printf("Input: %d, Reversed: %d\n", test, result)
	}
}
