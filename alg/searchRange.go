package alg

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

// 如果数组中不存在目标值 target，返回 [-1, -1]。

// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]

func SearchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		mid := (right-left)/2 + left

		if nums[mid] == target {
			res = []int{mid, mid}
			for i := 1; mid-i >= 0 && nums[mid-i] == target; i++ {
				res[0] = mid - i
			}
			for i := 1; mid+i < n && nums[mid+i] == target; i++ {
				res[1] = mid + i
			}
			break

		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return res
}
