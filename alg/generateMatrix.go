package alg

func GenerateMatrix(n int) [][]int {
	res := make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	i := 1

	up, down, left, right := 0, n-1, 0, n-1

	for up <= down && left <= right {
		for j := left; j <= right; j++ {
			res[up][j] = i
			i++
		}
		up++
		for j := up; j <= down; j++ {
			res[j][right] = i
			i++
		}
		right--

		if left > right || up > down {
			break
		}

		for j := right; j >= left; j-- {
			res[down][j] = i
			i++
		}
		down--
		for j := down; j >= up; j-- {
			res[j][left] = i
			i++
		}
		left++

	}

	return res
}
