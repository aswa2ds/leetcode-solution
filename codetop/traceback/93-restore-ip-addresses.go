package traceback

import "fmt"

func restoreIpAddresses(s string) []string {

	// 用于判断传入字符串是否满足一个IP 段
	check := func(s string) bool {
		// 不为空，且不超过3 位
		if len(s) == 0 || len(s) > 3 {
			return false
		}
		// 1位时，可以为0-9 任意数字
		if len(s) == 1 {
			return true
		}
		// 2-3 位时
		if len(s) != 1 {
			// 不可以0开头
			if s[0] == '0' {
				return false
			}
			// 不可大于255
			val := 0
			for i := 0; i < len(s); i++ {
				val = 10*val + int(s[i]-'0')
			}
			if val <= 255 {
				return true
			}
		}
		return false
	}

	var restoreIpAddressesOfSegs func(string, int) []string
	restoreIpAddressesOfSegs = func(s string, seg int) []string {
		// 递归终止条件：仅剩一段时，全部进行检查
		if seg == 1 {
			if check(s) {
				return []string{s}
			} else {
				return []string{}
			}
		}
		var res []string
		for i := 1; i <= len(s); i++ {
			s1 := s[:i]
			// 检查当前s1 是否可以作为IP 地址的一段
			if check(s1) {
				// 检查后续字符串是否可以分成seg-1 段
				subIpSegs := restoreIpAddressesOfSegs(s[i:], seg-1)
				// 不满足，说明当前段不可为s1
				if len(subIpSegs) == 0 {
					continue
				}
				// 满足，组装
				for _, subIpSeg := range subIpSegs {
					res = append(res, fmt.Sprintf("%s.%s", s1, subIpSeg))
				}
			} else {
				// 当前s2 不可作为IP 地址的一段，再增加位数也不可以满足，结束循环
				break
			}
		}
		return res
	}
	return restoreIpAddressesOfSegs(s, 4)
}
