package doublepoint

func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	for i < len(version1) || j < len(version2) {
		x := 0
		for ; i < len(version1) && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		i++
		y := 0
		for ; j < len(version2) && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j]-'0')
		}
		j++
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}

// 写得巨几把复杂，递归屎一样
//unc compareVersion(version1 string, version2 string) int {
//	v1Arr := strings.Split(version1, ".")
//	v2Arr := strings.Split(version2, ".")
//
//	var compare func(v1, v2 []string) int
//
//	compare = func(v1, v2 []string) int {
//		if len(v1) == 0 && len(v2) == 0 {
//			return 0
//		}
//
//		if len(v1) == 0 {
//			num, _ := strconv.Atoi(v2[0])
//			if num > 0 {
//				return -1
//			} else {
//				return compare(v1, v2[1:])
//			}
//		}
//
//		if len(v2) == 0 {
//			num, _ := strconv.Atoi(v1[0])
//			if num > 0 {
//				return 1
//			} else {
//				return compare(v1[1:], v2)
//			}
//		}
//
//		num1, _ := strconv.Atoi(v1[0])
//		num2, _ := strconv.Atoi(v2[0])
//
//		if num1 > num2 {
//			return 1
//		}
//		if num1 < num2 {
//			return -1
//		}
//		return compare(v1[1:], v2[1:])
//	}
//
//	return compare(v1Arr, v2Arr)
//}
