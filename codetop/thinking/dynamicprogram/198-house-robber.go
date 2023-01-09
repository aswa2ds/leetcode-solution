package dynamicprogram

func rob(nums []int) int {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	dp0, dp1 := 0, 0
	for _, num := range nums {
		dp0, dp1 = dp1, max(dp0+num, dp1)
	}

	return dp1
}

// 莫名其妙写的巨麻烦
//func rob(nums []int) int {
//	max := func(a, b int) int {
//		if a > b {
//			return a
//		}
//		return b
//	}
//
//	n := len(nums)
//
//	if n == 1 {
//		return nums[0]
//	}
//	if n == 2 {
//		return max(nums[0], nums[1])
//	}
//	if n == 3 {
//		return max(nums[0]+nums[2], nums[1])
//	}
//
//	dp0, dp1, dp2 := nums[0], max(nums[0], nums[1]), max(nums[0]+nums[2], nums[1])
//
//	for i := 3; i < n; i++ {
//		dp0, dp1, dp2 = dp1, dp2, max(dp0, dp1)+nums[i]
//	}
//
//	return max(dp1, dp2)
//}
