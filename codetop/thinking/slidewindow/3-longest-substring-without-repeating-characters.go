package slidewindow

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	chMap := make(map[uint8]int)
	maxLen := 0
	for right < len(s) {
		ch := s[right]
		if idx, ok := chMap[ch]; ok && idx != -1 {
			for left <= idx {
				chMap[s[left]] = -1
				left++
			}
		}
		chMap[s[right]] = right
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
		right++
	}
	return maxLen
}
