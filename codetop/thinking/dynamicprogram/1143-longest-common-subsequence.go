package dynamicprogram

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text2) > len(text1) {
		text1, text2 = text2, text1
	}

	dp := make([]int, len(text2)+1)
	leftTop := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for row := range text1 {
		leftTop, dp[0] = dp[0], 0

		for col := range text2 {
			if text1[row] == text2[col] {
				leftTop, dp[col+1] = dp[col+1], leftTop+1
			} else {
				leftTop, dp[col+1] = dp[col+1], max(dp[col], dp[col+1])
			}
		}
	}

	return dp[len(dp)-1]
}
