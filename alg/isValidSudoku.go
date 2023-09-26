package alg

func isValidSudoku(board [][]byte) bool {
	var rows, cols [9][9]int
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, col := range row {
			if col == '.' {
				continue
			}
			index := col - '1'
			rows[i][index]++
			cols[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || cols[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true

}
