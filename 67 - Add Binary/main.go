package main

func addBinary(a string, b string) string {
	result := ""
	carry := 0
	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 || carry > 0 {
		digitA := 0
		if i >= 0 {
			digitA = int(a[i] - '0')
			i--
		}

		digitB := 0
		if j >= 0 {
			digitB = int(b[j] - '0')
			j--
		}

		sum := digitA + digitB + carry
		result = string(rune('0'+sum%2)) + result
		carry = sum / 2
	}

	return result
}
