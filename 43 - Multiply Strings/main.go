package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1, len2 := len(num1), len(num2)
	result := make([]int, len1+len2)

	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			mul := (int(num1[i]) - '0') * (int(num2[j]) - '0')
			pos1, pos2 := i+j, i+j+1
			sum := mul + result[pos2]

			result[pos2] = sum % 10
			result[pos1] += sum / 10
		}
	}

	res := ""
	for _, digit := range result {
		if !(res == "" && digit == 0) {
			res += string(rune(digit + '0'))
		}
	}

	return res
}
