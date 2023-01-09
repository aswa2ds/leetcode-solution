package array

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	top, leader, bottom, trailing := 0, 0, m-1, n-1
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	i, j := 0, 0
	direction := 0
	res := []int{matrix[0][0]}
	for len(res) < m*n {
		newI, newJ := i+directions[direction][0], j+directions[direction][1]
		if newJ > trailing {
			top++
			direction = (direction + 1) % 4
			continue
		}
		if newI > bottom {
			trailing--
			direction = (direction + 1) % 4
			continue
		}
		if newJ < leader {
			bottom--
			direction = (direction + 1) % 4
			continue
		}
		if newI < top {
			leader++
			direction = (direction + 1) % 4
			continue
		}

		res = append(res, matrix[newI][newJ])
		i, j = newI, newJ
	}

	return res
}
