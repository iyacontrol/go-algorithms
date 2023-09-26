package alg

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {

	res := [][]int{}

	sort.Ints(candidates)

	var dfs func(start int, temp []int, sum int)
	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				newTemp := make([]int, len(temp))
				copy(newTemp, temp)
				res = append(res, newTemp)
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			if i-1 >= start && candidates[i-1] == candidates[i] {
				continue
			}

			temp = append(temp, candidates[i])
			dfs(i+1, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0, []int{}, 0)
	return res
}
