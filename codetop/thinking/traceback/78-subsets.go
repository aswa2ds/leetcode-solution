package traceback

func subsets(nums []int) [][]int {
	res := [][]int{{}}

	for _, num := range nums {
		n := len(res)
		for i := 0; i < n; i++ {
			res = append(res, append([]int{num}, res[i]...))
		}
	}

	return res
}
