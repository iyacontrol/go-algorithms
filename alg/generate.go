package alg

func Generate(numRows int) [][]int {
	res := make([][]int, numRows)

	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = res[i-1][j-1] + res[i-1][j]
			}
		}

		res[i] = row
	}

	return res
}
