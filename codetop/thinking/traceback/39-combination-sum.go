package traceback

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	for i, candidate := range candidates {
		if candidate == target {
			res = append(res, []int{candidate})
		}
		if candidate < target {
			subCombSums := combinationSum(candidates[i:], target-candidate)
			if len(subCombSums) != 0 {
				for _, subCombSum := range subCombSums {
					res = append(res, append(subCombSum, candidate))
				}
			}
		}
	}
	return res
}
