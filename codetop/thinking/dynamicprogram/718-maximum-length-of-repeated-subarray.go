package dynamicprogram

func findLength(nums1 []int, nums2 []int) int {
	dp := make([]int, len(nums1)+1)
	leftTop := 0

	res := 0

	for row := range nums1 {
		leftTop = 0
		for col := range nums2 {
			if nums1[row] == nums2[col] {
				leftTop, dp[col+1] = dp[col+1], leftTop+1
			} else {
				leftTop, dp[col+1] = dp[col+1], 0
			}
			if dp[col+1] > res {
				res = dp[col+1]
			}
		}
	}

	return res
}
