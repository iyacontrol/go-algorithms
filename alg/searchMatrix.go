package alg

func SearchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	left, right := 0, m*n-1

	for left <= right {
		mid := (right-left)/2 + left

		row := mid / n
		col := mid % n

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false

}
