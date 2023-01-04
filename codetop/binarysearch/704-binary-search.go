package binarysearch

func searchI(nums []int, target int) int {
	n := len(nums)

	lo, hi := 0, n-1

	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			lo, hi = mid, mid
			break
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if nums[lo] == target {
		return lo
	}

	return -1
}
