package dynamicprogram

import "math"

func minPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0])+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for _, line := range grid {
		for i, num := range line {
			dp[i+1] = min(dp[i+1], dp[i]) + num
		}
		dp[0] = math.MaxInt
	}

	return dp[len(dp)-1]
}
