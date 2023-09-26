package alg

func SpiderMatrix(m, n int) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	res[0][0] = 1

	// 枚举，direc=0 表示走向是横向右, direc = 1 表示竖向下, direc=2 表示对角线向上 ，direc=3 表示对角线向下
	direc := 0
	row, col := 0, 0

	for i := 2; i <= m*n; i++ {

		if direc == 0 {
			col++
			if row == 0 {
				direc = 3
			} else if row == m-1 {
				direc = 2
			}
		} else if direc == 1 {
			row++
			if col == 0 {
				direc = 2
			} else if col == n-1 {
				direc = 3
			}
		} else if direc == 2 {
			col++
			row--
			if row == 0 && col != n-1 {
				direc = 0
			} else if col == n-1 {
				direc = 1
			}
		} else {
			col--
			row++
			if col == 0 && row != m-1 {
				direc = 1
			} else if row == m-1 {
				direc = 0
			}
		}

		res[row][col] = i

	}

	return res
}
