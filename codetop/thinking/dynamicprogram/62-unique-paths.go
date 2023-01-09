package dynamicprogram

import "math/big"

func uniquePaths(m int, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

// 动态规划
//func uniquePaths(m int, n int) int {
//	dp := make([]int, n+1)
//	dp[0] = 1
//
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			dp[j+1] = dp[j] + dp[j+1]
//		}
//		dp[0] = 0
//	}
//
//	return dp[len(dp)-1]
//}
