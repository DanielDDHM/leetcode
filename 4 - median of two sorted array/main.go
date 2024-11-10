package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Certifique-se de que nums1 é o array menor para otimizar a busca binária
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxX := -1 << 63 // equivalente a math.MinInt64
		if partitionX > 0 {
			maxX = nums1[partitionX-1]
		}
		minX := 1<<63 - 1 // equivalente a math.MaxInt64
		if partitionX < x {
			minX = nums1[partitionX]
		}

		maxY := -1 << 63
		if partitionY > 0 {
			maxY = nums2[partitionY-1]
		}
		minY := 1<<63 - 1
		if partitionY < y {
			minY = nums2[partitionY]
		}

		// Verifica se encontrou a partição correta
		if maxX <= minY && maxY <= minX {
			// Verifica se a quantidade total de elementos é par ou ímpar
			if (x+y)%2 == 0 {
				return float64(max(maxX, maxY)+min(minX, minY)) / 2
			} else {
				return float64(max(maxX, maxY))
			}
		} else if maxX > minY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	panic("Input arrays are not sorted or have invalid lengths")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // Output: 2.0
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // Output: 2.5
}
