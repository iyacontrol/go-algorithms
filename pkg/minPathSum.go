package pkg

import "math"

func MinPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := math.MaxInt

	var dfs func(row, col, sum int)
	dfs = func(row, col, sum int) {
		if row == m-1 && col == n-1 {
			res = min(res, sum)
			return
		}

		// 两种
		if col < n-1 {
			dfs(row, col+1, sum+grid[row][col+1])
		}

		if row < m-1 {
			dfs(row+1, col, sum+grid[row+1][col])
		}

	}

	dfs(0, 0, grid[0][0])

	if res == math.MaxInt {
		return 0
	}

	return res
}
