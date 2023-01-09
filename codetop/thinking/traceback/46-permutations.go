package traceback

// 空间复杂度较高
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := make([][]int, 0)
	for i, num := range nums {
		subNums := append(make([]int, 0), nums[:i]...)
		subNums = append(subNums, nums[i+1:]...)
		subPs := permute(subNums)
		for _, subP := range subPs {
			res = append(res, append([]int{num}, subP...))
		}
	}
	return res
}
