package pkg

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

// 示例 1：

// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

func Permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)

	visited := map[int]bool{}

	var dfs func(temp []int)

	dfs = func(temp []int) {
		if len(temp) == n {
			newTemp := make([]int, len(temp))
			copy(newTemp, temp)
			res = append(res, newTemp)
			return
		}

		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}

			temp = append(temp, nums[i])
			visited[i] = true
			dfs(temp)
			temp = temp[:len(temp)-1]
			visited[i] = false
		}
	}

	dfs([]int{})

	return res
}
