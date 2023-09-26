package alg

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// 示例 1：

// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

func Subsets(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)

	var dfs func(index int, temp []int)
	dfs = func(index int, temp []int) {
		// 终止条件
		if index > n-1 {
			newTemp := make([]int, len(temp))
			copy(newTemp, temp)
			res = append(res, newTemp)

			return
		}

		// 选择
		temp = append(temp, nums[index])
		dfs(index+1, temp)
		temp = temp[:len(temp)-1]

		// not
		dfs(index+1, temp)

	}

	dfs(0, []int{})

	return res
}
