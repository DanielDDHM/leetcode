package main

func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}

	negative := (dividend < 0) != (divisor < 0)

	a := int64(dividend)
	if a < 0 {
		a = -a
	}

	b := int64(divisor)
	if b < 0 {
		b = -b
	}

	var result int64
	for a >= b {
		var shift int64
		for (b << (shift + 1)) <= a {
			shift++
		}
		result += 1 << shift
		a -= b << shift
	}

	if negative {
		result = -result
	}

	if result > 2147483647 {
		return 2147483647
	}
	if result < -2147483648 {
		return -2147483648
	}

	return int(result)
}
