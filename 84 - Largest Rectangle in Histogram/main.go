package main

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
