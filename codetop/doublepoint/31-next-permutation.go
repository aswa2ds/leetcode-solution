package doublepoint

func nextPermutation(nums []int) {
	length := len(nums)

	var j int
	for j = length - 1; j > 0; j-- {
		i := j - 1
		if nums[i] < nums[j] {
			for k := length - 1; k >= j; k-- {
				if nums[k] > nums[i] {
					nums[i], nums[k] = nums[k], nums[i]
					break
				}
			}
			break
		}
	}

	func(i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}(j, length-1)
}

// 时间空间复杂度没毛病，稍作代码逻辑优化
//func nextPermutation(nums []int) {
//	length := len(nums)
//	if length == 0 || length == 1 {
//		return
//	}
//
//	reverse := func(i, j int) {
//		for i < j {
//			nums[i], nums[j] = nums[j], nums[i]
//			i++
//			j--
//		}
//	}
//
//	found := false
//	for j := length - 1; j >= 1; j-- {
//		i := j - 1
//		if nums[i] < nums[j] {
//			lastBiggerThanI := j
//			for lastBiggerThanI < length {
//				if nums[lastBiggerThanI] <= nums[i] {
//					break
//				}
//				lastBiggerThanI++
//			}
//			lastBiggerThanI--
//			nums[i], nums[lastBiggerThanI] = nums[lastBiggerThanI], nums[i]
//			reverse(j, length-1)
//			found = true
//			return
//		}
//	}
//	if !found {
//		reverse(0, length-1)
//	}
//}
