package traceback

func generateParenthesis(n int) []string {

	var res []string

	var traceback func(n, sub int, base string)

	traceback = func(n, sub int, base string) {
		// 生成完成
		if n == 0 {
			res = append(res, base)
			return
		}
		// 如果sub 为0，说明当前情况左右括号数量相同，只能生成一个左括号
		if sub == 0 {
			traceback(n-1, sub+1, base+"(")
			return
		}
		// 如果sub 和n 相同，说明当前剩余的位置只能为右括号
		if sub == n {
			traceback(n-1, sub-1, base+")")
			return
		}
		// sub < n 且 sub != 0 的情况，下一个括号可以随意加
		traceback(n-1, sub+1, base+"(")
		traceback(n-1, sub-1, base+")")
	}

	traceback(2*n, 0, "")

	return res

}
