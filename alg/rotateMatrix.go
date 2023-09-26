package alg

func RotateMatrix(matrix [][]int) {
	n := len(matrix)

	// 圈循环，i表示一圈， i 等于0 表示最外侧的一圈
	for i := 0; i < n/2; i++ {
		up, down, left, right := i, n-1-i, i, n-1-i

		// 具体的一圈中，j需要旋转的数字

		for j := 0; j < n-2*i-1; j++ {
			matrix[up][left+j], matrix[up+j][right], matrix[down][right-j], matrix[down-j][left] = matrix[down-j][left], matrix[up][left+j], matrix[up+j][right], matrix[down][right-j]
		}

	}

}
