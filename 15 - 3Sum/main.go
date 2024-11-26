package main

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	n := len(nums)

	// Ordenar os números para simplificar a busca
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// Evitar duplicatas para o primeiro número
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Evitar duplicatas para o segundo e terceiro números
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
