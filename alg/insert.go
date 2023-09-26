package alg

import "sort"

func Insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	n := len(intervals)

	for i := 0; i < n; i++ {
		if intervals[i][0] > newInterval[1] || intervals[i][1] < newInterval[0] {
			res = append(res, intervals[i])
		} else {
			minLeftEdge := min(intervals[i][0], newInterval[0])
			maxRightEdge := max(intervals[i][1], newInterval[1])
			newInterval[0] = minLeftEdge
			newInterval[1] = maxRightEdge
		}
	}

	res = append(res, newInterval)

	sort.Slice(res, func(i, j int) bool {
		if res[i][0] < res[j][0] {
			return true
		}
		return false
	})

	return res
}
