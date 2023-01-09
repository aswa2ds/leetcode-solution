package dynamicprogram

import "math"

func maxSubArray(nums []int) int {
	pre := 0
	res := math.MinInt

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, num := range nums {
		pre = max(pre+num, num)
		res = max(res, pre)
	}

	return res
}
