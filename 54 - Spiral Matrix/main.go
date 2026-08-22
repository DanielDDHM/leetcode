package main

func spiralOrder(matrix [][]int) (ans []int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			ans = append(ans, matrix[top][col])
		}
		top++

		for row := top; row <= bottom; row++ {
			ans = append(ans, matrix[row][right])
		}
		right--

		if top <= bottom {
			for col := right; col >= left; col-- {
				ans = append(ans, matrix[bottom][col])
			}
			bottom--
		}

		if left <= right {
			for row := bottom; row >= top; row-- {
				ans = append(ans, matrix[row][left])
			}
			left++
		}
	}

	return
}
