package dynamicprogram

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp0 := 1
	dp1 := 2

	for i := 3; i <= n; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}

	return dp1
}
