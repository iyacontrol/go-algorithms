package alg

// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
// 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：[[1,0,1],[0,0,0],[1,0,1]]

func SetZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	setZero := func(row, col int) {
		matrix[row][col] = 0
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == 0 {
				// 元素所在的行列标记，后续置零
				// 单行
				for i := 0; i < n; i++ {
					defer setZero(row, i)
				}
				// 单列
				for j := 0; j < m; j++ {
					defer setZero(j, col)
				}
			}
		}
	}

}
