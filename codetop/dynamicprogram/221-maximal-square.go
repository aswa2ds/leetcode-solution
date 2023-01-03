package dynamicprogram

func maximalSquare(matrix [][]byte) int {
	leftTop := 0
	dp := make([]int, len(matrix[0])+1)

	min := func(a, b, c int) int {
		if a > b {
			a, b = b, a
		}
		if a < c {
			return a
		}
		return c
	}

	maxSide := 0

	for _, line := range matrix {
		leftTop = 0
		for i, num := range line {
			if num == '1' {
				leftTop, dp[i+1] = dp[i+1], min(leftTop, dp[i], dp[i+1])+1
				if dp[i+1] > maxSide {
					maxSide = dp[i+1]
				}
			} else {
				leftTop, dp[i+1] = dp[i+1], 0
			}
		}
	}

	return maxSide * maxSide
}
