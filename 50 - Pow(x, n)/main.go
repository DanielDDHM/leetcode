package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	if n%2 == 0 {
		half := myPow(x, n/2)
		return half * half
	}

	return x * myPow(x, n-1)
}
