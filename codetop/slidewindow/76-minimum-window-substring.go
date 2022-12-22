package slidewindow

import "math"

func minWindow(s string, t string) string {
	// 用于计数剩余需要匹配的字符数量，初始为t 的长度
	leftNumToMatch := len(t)
	// 桶计数的方式，可能出现的字母一共只有52个，因此创建一个容量为60的数组
	letterRecords := make([]int, 60)
	// 记录当前匹配到满足条件的子串的最小长度，初始为最大Int
	minLen := math.MaxInt
	// 记录最终答案，如果没有匹配到结果则返回“”，因此初始值为“”
	res := ""

	// 首先将t 中的字符计数统计到桶中
	for i := 0; i < len(t); i++ {
		letterRecords[t[i]-'A']++
	}

	// 双指针，滑动窗口的方式进行遍历
	for left, right := 0, 0; right < len(s); right++ {
		// 对于右指针的字符，首先将对应桶中的计数减一，此时抽象两种情况：
		// 1. s[left:right] 中的该字符数量小于t 中的该字符数量（t 中存在的字符），那么桶中本轮循环的初始值大于0，计数减一后，桶中数量仍大于等于0
		// 2. s[left:right] 中的该字符数量大于等于t 中该字符的数量(t 中不存在的字符 或 t 中已经记够数量的字符），那么桶中本轮循环的初始值小于等于0，计数减一后，桶中数量小于0
		// 对于一种情况，说明当前该字符的匹配量不够，需要减小leftNumToMatch 的计数；第二种情况，说明该字符的匹配量够了，不需要减小leftNumToMatch 的计数
		letterRecords[s[right]-'A']--
		if letterRecords[s[right]-'A'] >= 0 {
			// leftNumToMatch 不会出现增加的情况，因为只有保证leftNumToMatch 不会增加的情况，才会进行左指针的右移
			leftNumToMatch--
		}

		// 对于左指针，当左指针的字符对应桶中的计数小于0 时，说明，当前s[left:right+1] 中所包含的该字符，超出匹配t 的需求，可以将该字符剔除，并且不影响leftNumToMatch 计数
		for left <= right && letterRecords[s[left]-'A'] < 0 {
			letterRecords[s[left]-'A']++
			left++
		}

		// 如果此时leftNumToMatch 已经达到0，说明s[left:right+1] 已经覆盖了t，那么只需要比较当前子串长度和已经记录的最小长度之间的大小关系，记录最小子串即可
		if leftNumToMatch == 0 && right-left+1 < minLen {
			minLen = right - left + 1
			res = s[left : right+1]
		}

	}

	return res
}

//func minWindow(s string, t string) string {
//	tChMap, currMap := make(map[byte]int), make(map[byte]int)
//
//	minMatchLen := math.MaxInt
//	ansL, ansR := -1, -1
//
//	for i := 0; i < len(t); i++ {
//		tChMap[t[i]]++
//	}
//
//	check := func() bool {
//		for k, v := range tChMap {
//			if currMap[k] < v {
//				return false
//			}
//		}
//		return true
//	}
//
//	for l, r := 0, 0; r < len(s); r++ {
//		if tChMap[s[r]] > 0 {
//			currMap[s[r]]++
//		}
//		for check() && l <= r {
//			if r-l+1 < minMatchLen {
//				minMatchLen = r - l + 1
//				ansL, ansR = l, l+minMatchLen
//			}
//			if _, ok := tChMap[s[l]]; ok {
//				currMap[s[l]]--
//			}
//			l++
//		}
//	}
//
//	if ansL == -1 {
//		return ""
//	}
//
//	return s[ansL:ansR]
//
//}
