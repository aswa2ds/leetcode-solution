package dynamicprogram

func minDistance(word1 string, word2 string) int {
	cols := len(word2) + 1
	dp := make([]int, cols)
	lefttop := 0
	for i := range dp {
		dp[i] = i
	}
	minOf3 := func(a, b, c int) int {
		if a > b {
			a, b = b, a
		}
		if a < c {
			return a
		}
		return c
	}

	for row := range word1 {
		lefttop = dp[0]
		dp[0] = row + 1
		for col := range word2 {
			if word1[row] == word2[col] {
				tmp := lefttop
				lefttop = dp[col+1]
				dp[col+1] = tmp
			} else {
				tmp := minOf3(dp[col], lefttop, dp[col+1]) + 1
				lefttop = dp[col+1]
				dp[col+1] = tmp
			}
		}
	}

	return dp[len(dp)-1]
}
