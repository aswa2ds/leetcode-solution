package dynamicprogram

func maxProfit(prices []int) int {
	profile := 0

	// 购买次数不限，只要价格高了，就直接卖出去，毕竟买入不要钱的
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profile += prices[i] - prices[i-1]
		}
	}

	return profile
}
