package binarysearch

func search(nums []int, target int) int {
	n := len(nums)

	lo, hi := 0, n-1

	for lo <= hi {
		mid := (lo + hi) >> 1
		if target == nums[mid] {
			return mid
		}

		if nums[mid] >= nums[lo] {
			if target < nums[mid] && target >= nums[lo] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if target > nums[mid] && target < nums[lo] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}
