package alg

import "fmt"

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 请必须使用时间复杂度为 O(log n) 的算法。

// 示例 1:

// 输入: nums = [1,3,5,6], target = 4
// 输出: 2

func SearchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if left == right {
				fmt.Println("如果走到这里，则说明目标值比重合处值小，因为是插入，重合处就是需要插入的位置")
			}

			right = mid - 1
		} else {
			//
			if left == right {
				fmt.Println("如果走到这里，则说明目标值比重合处值大，所以需要+1")
			}

			left = mid + 1
		}

	}

	return left

}
