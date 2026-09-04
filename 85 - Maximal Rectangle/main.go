package main

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])
	heights := make([]int, n)
	maxArea := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		area := largestRectangleArea(heights)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			var w int
			if len(stack) == 0 {
				w = i
			} else {
				w = i - stack[len(stack)-1] - 1
			}

			area := h * w
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		h := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]

		var w int
		if len(stack) == 0 {
			w = len(heights)
		} else {
			w = len(heights) - stack[len(stack)-1] - 1
		}

		area := h * w
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
