package main

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	if k >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	buy := make([]int, k+1)
	sell := make([]int, k+1)

	for i := 0; i <= k; i++ {
		buy[i] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for j := k; j >= 1; j-- {
			sell[j] = max(sell[j], buy[j]+prices[i])
			buy[j] = max(buy[j], sell[j-1]-prices[i])
		}
	}

	return sell[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
