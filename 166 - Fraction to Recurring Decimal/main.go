package main

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	var result string

	if (numerator < 0) != (denominator < 0) && numerator != 0 {
		result += "-"
	}

	num := int64(numerator)
	den := int64(denominator)
	if num < 0 {
		num = -num
	}
	if den < 0 {
		den = -den
	}

	result += strconv.FormatInt(num/den, 10)
	remainder := num % den
	if remainder == 0 {
		return result
	}

	result += "."
	remainderMap := make(map[int64]int)

	for remainder != 0 {
		if pos, seen := remainderMap[remainder]; seen {
			result = result[:pos] + "(" + result[pos:] + ")"
			return result
		}

		remainderMap[remainder] = len(result)
		remainder *= 10
		result += strconv.FormatInt(remainder/den, 10)
		remainder %= den
	}

	return result
}
