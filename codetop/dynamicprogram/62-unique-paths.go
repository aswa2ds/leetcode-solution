package dynamicprogram

func uniquePaths(m int, n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[j+1] = dp[j] + dp[j+1]
		}
		dp[0] = 0
	}

	return dp[len(dp)-1]
}
