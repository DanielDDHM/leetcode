package main

import (
	"fmt"
	"math"
)

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	maxCount := 0

	for i := 0; i < len(points); i++ {
		slopes := make(map[string]int)
		duplicates := 0
		currentMax := 0

		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

			dy := points[j][1] - points[i][1]
			dx := points[j][0] - points[i][0]

			if dy == 0 && dx == 0 {
				duplicates++
				continue
			}

			gcd := gcdFunc(dy, dx)
			dy /= gcd
			dx /= gcd

			if dx < 0 {
				dx = -dx
				dy = -dy
			}

			key := fmt.Sprintf("%d/%d", dy, dx)
			slopes[key]++
			if slopes[key] > currentMax {
				currentMax = slopes[key]
			}
		}

		if currentMax+duplicates+1 > maxCount {
			maxCount = currentMax + duplicates + 1
		}
	}

	return maxCount
}

func gcdFunc(a, b int) int {
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))

	for b != 0 {
		a, b = b, a%b
	}

	if a == 0 {
		return 1
	}
	return a
}
