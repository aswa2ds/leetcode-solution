package slidewindow

// 这个问题有一个点，那就是，后进入滑动窗口中的较大数，会比在它之前的较小数后出窗口
func maxSlidingWindow(nums []int, k int) []int {
	deque := make([]int, 0)

	push := func(index int) {
		for len(deque) > 0 && nums[index] > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, index)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	res := make([]int, 0)
	res = append(res, nums[deque[0]])
	for i := k; i < len(nums); i++ {
		push(i)
		if deque[0] <= i-k {
			deque = deque[1:]
		}
		res = append(res, nums[deque[0]])
	}
	return res
}

// 优先级队列 事件复杂度过高，歇逼
//func maxSlidingWindow(nums []int, k int) []int {
//	maxHeap := make([]int, k)
//	for i := 0; i < k; i++ {
//		maxHeap[i] = math.MinInt
//	}
//
//	var shiftDown func(int)
//	var shiftUp func(int)
//
//	shiftDown = func(index int) {
//		maxValIndex := index
//		leftLeafIndex := 2*index + 1
//		if leftLeafIndex < k && maxHeap[index] < maxHeap[leftLeafIndex] {
//			maxValIndex = leftLeafIndex
//		}
//
//		rightLeafIndex := leftLeafIndex + 1
//		if rightLeafIndex < k && maxHeap[maxValIndex] < maxHeap[rightLeafIndex] {
//			maxValIndex = rightLeafIndex
//		}
//
//		if maxValIndex > index {
//			maxHeap[index], maxHeap[maxValIndex] = maxHeap[maxValIndex], maxHeap[index]
//			shiftDown(maxValIndex)
//		}
//	}
//
//	shiftUp = func(index int) {
//		if index == 0 {
//			return
//		}
//		rootIndex := (index - 1) / 2
//		if maxHeap[rootIndex] < maxHeap[index] {
//			maxHeap[rootIndex], maxHeap[index] = maxHeap[index], maxHeap[rootIndex]
//			shiftUp(rootIndex)
//		}
//	}
//
//	push := func(newVal int, oldVal int) {
//		if oldVal == newVal {
//			return
//		}
//		for i := 0; i < k; i++ {
//			if maxHeap[i] == oldVal {
//				maxHeap[i] = newVal
//				if newVal < oldVal {
//					shiftDown(i)
//				} else {
//					shiftUp(i)
//				}
//				break
//			}
//		}
//	}
//
//	res := make([]int, 0)
//	numsWithHeads := make([]int, 0)
//	for i := 0; i < k; i++ {
//		numsWithHeads = append(numsWithHeads, math.MinInt)
//	}
//	numsWithHeads = append(numsWithHeads, nums...)
//
//	for left, right := 0, k; right < len(numsWithHeads); left, right = left+1, right+1 {
//		push(numsWithHeads[right], numsWithHeads[left])
//		res = append(res, maxHeap[0])
//	}
//
//	return res[k-1:]
//}
