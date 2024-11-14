package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1 // Inicializa os dois ponteiros
	maxArea := 0                    // Armazena a área máxima encontrada

	for left < right {
		// Calcula a área atual entre os ponteiros esquerdo e direito
		minHeight := min(height[left], height[right])

		area := minHeight * (right - left)

		// Atualiza a área máxima, se a área atual for maior
		if maxArea < area {
			maxArea = area
		}

		// Move o ponteiro que está no lado menor para tentar encontrar uma área maior
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
