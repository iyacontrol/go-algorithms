package alg

import "sort"

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

// 你返回所有和为 0 且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	// 先排序
	sort.Ints(nums)

	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				right--
				left++
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}

	}

	return res
}
