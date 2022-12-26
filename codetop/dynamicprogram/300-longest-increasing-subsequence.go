package dynamicprogram

func lengthOfLIS(nums []int) int {
	var dp []int

	dp = append(dp, nums[0])
	for _, num := range nums {
		if num > dp[len(dp)-1] {
			dp = append(dp, num)
		} else {
			lo, hi := 0, len(dp)-1
			for lo < hi {
				mid := (lo + hi) / 2
				if dp[mid] < num {
					lo = mid + 1
				} else {
					hi = mid
				}
			}
			dp[lo] = num
		}
	}
	return len(dp)
}

//func lengthOfLIS(nums []int) int {
//	dp := make([]int, len(nums))
//
//	dp[0] = 1
//	maxLen := 1
//
//	for i := 1; i < len(nums); i++ {
//		tmpMax := 0
//		for j := i - 1; j >= 0; j-- {
//			if nums[j] < nums[i] {
//				if dp[j] > tmpMax {
//					tmpMax = dp[j]
//				}
//			}
//		}
//		dp[i] = tmpMax + 1
//		if dp[i] > maxLen {
//			maxLen = dp[i]
//		}
//	}
//
//	return maxLen
//}
