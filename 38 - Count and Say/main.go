package main

func countAndSay(n int) string {
	result := "1"
	for i := 1; i < n; i++ {
		result = encode(result)
	}
	return result
}

func encode(s string) string {
	var ans []byte
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		count := j - i
		ans = append(ans, byte('0'+count))
		ans = append(ans, s[i])
		i = j
	}
	return string(ans)
}
