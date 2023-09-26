package alg

func SpiralOrder(matrix [][]int) []int {
	res := []int{}
	m := len(matrix)
	n := len(matrix[0])

	left, up, right, down := 0, 0, n-1, m-1

	for left <= right && up <= down {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++

		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if up > down || left > right {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--

		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res

}
