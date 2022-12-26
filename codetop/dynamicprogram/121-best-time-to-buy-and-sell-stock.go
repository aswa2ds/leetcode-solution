package dynamicprogram

func maxProfit_i(prices []int) int {
	maxP := 0
	buy := prices[0]

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, price := range prices {
		if price > buy {
			maxP = max(price-buy, maxP)
		} else {
			buy = price
		}
	}

	return maxP
}
