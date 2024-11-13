package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}

func main() {
	testCases := []int{121, -121, 10, 12321, 123}

	// isPalindrome(121) = true
	// isPalindrome(-121) = false
	// isPalindrome(10) = false
	// isPalindrome(12321) = true
	// isPalindrome(123) = false

	for _, num := range testCases {
		fmt.Printf("isPalindrome(%d) = %v\n", num, isPalindrome(num))
	}
}
