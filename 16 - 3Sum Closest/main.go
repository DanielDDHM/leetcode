package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	// Ordenar os números para simplificar a busca
	sort.Ints(nums)

	n := len(nums)
	closestSum := nums[0] + nums[1] + nums[2] // Inicializar com a primeira soma possível

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// Atualizar a soma mais próxima, se necessário
			if abs(target-sum) < abs(target-closestSum) {
				closestSum = sum
			}

			// Mover os ponteiros com base na soma
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				// Se a soma for exatamente igual ao target, retorne diretamente
				return sum
			}
		}
	}

	return closestSum
}

// Função auxiliar para calcular o valor absoluto
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
