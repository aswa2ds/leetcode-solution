package sortalgorithm

import (
	"math/rand"
	"time"
)

func merge(intervals [][]int) [][]int {
	rand.Seed(time.Now().UnixNano())
	partition := func(left, right int) int {
		mark := left
		key := rand.Intn(right-left) + left
		pivot := intervals[key]
		intervals[key], intervals[left] = intervals[left], intervals[key]
		for i := left + 1; i <= right; i++ {
			if intervals[i][0] < pivot[0] {
				mark++
				intervals[mark], intervals[i] = intervals[i], intervals[mark]
			}
		}
		intervals[left], intervals[mark] = intervals[mark], pivot
		return mark
	}

	var sort func(int, int)
	sort = func(lo, hi int) {
		if lo >= hi {
			return
		}
		pivotIndex := partition(lo, hi)
		sort(lo, pivotIndex-1)
		sort(pivotIndex+1, hi)
	}

	sort(0, len(intervals)-1)

	var res [][]int
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			if intervals[i][1] > res[len(res)-1][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}
