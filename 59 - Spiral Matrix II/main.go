package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	top, bottom, left, right := 0, n-1, 0, n-1
	num := 1

	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			matrix[top][col] = num
			num++
		}
		top++

		for row := top; row <= bottom; row++ {
			matrix[row][right] = num
			num++
		}
		right--

		if top <= bottom {
			for col := right; col >= left; col-- {
				matrix[bottom][col] = num
				num++
			}
			bottom--
		}

		if left <= right {
			for row := bottom; row >= top; row-- {
				matrix[row][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}
