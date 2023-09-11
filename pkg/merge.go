package pkg

import "sort"

func Merge(intervals [][]int) [][]int {

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
		} else {
			maxEdge := max(res[len(res)-1][1], intervals[i][1])
			res[len(res)-1][1] = maxEdge
		}
	}

	return res
}
