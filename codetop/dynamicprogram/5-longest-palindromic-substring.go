package dynamicprogram

import (
	"strings"
)

// Manacher 算法
func longestPalindrome(s string) string {
	preprocess := func(s string) string {
		n := len(s)
		if n == 0 {
			return "^$"
		}
		res := strings.Builder{}
		res.WriteRune('^')
		for _, ch := range s {
			res.WriteRune('#')
			res.WriteRune(ch)
		}
		res.WriteRune('#')
		res.WriteRune('$')
		return res.String()
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	t := preprocess(s)

	n := len(t)
	P := make([]int, n)
	C, R := 0, 0

	for i := 0; i < n; i++ {
		i_mirror := 2*C - i
		if R > i {
			P[i] = min(R-i, P[i_mirror])
		} else {
			P[i] = 0
		}

		for i-P[i]-1 >= 0 && i+P[i]+1 < n && t[i+P[i]+1] == t[i-P[i]-1] {
			P[i]++
		}

		if (i + P[i]) > R {
			C = i
			R = i + P[i]
		}
	}

	maxLen := 0
	centerIndex := 0

	for i := 2; i < n-2; i++ {
		if P[i] > maxLen {
			maxLen = P[i]
			centerIndex = i
		}
	}

	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}

//func longestPalindrome(s string) string {
//	if len(s) == 0 {
//		return s
//	}
//	res := ""
//	maxLen := 0
//	dp := make([]int, len(s))
//	for i := len(s) - 1; i >= 0; i-- {
//		for j := len(s) - 1; j >= i; j-- {
//			switch j - i {
//			case 0:
//				dp[j] = 1
//				if maxLen < 1 {
//					maxLen = 1
//					res = s[j : j+1]
//				}
//				break
//			case 1:
//				if s[i] == s[j] {
//					dp[j] = 2
//					if maxLen < 2 {
//						maxLen = 2
//						res = s[i : j+1]
//					}
//				} else {
//					dp[j] = 0
//				}
//				break
//			default:
//				if s[i] == s[j] && dp[j-1] != 0 {
//					dp[j] = dp[j-1] + 2
//					if j-i+1 > maxLen {
//						maxLen = j - i + 1
//						res = s[i : j+1]
//					}
//				} else {
//					dp[j] = 0
//				}
//			}
//		}
//	}
//	return res
//}
