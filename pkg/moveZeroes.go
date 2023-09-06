package pkg

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

// 示例 1:

// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]

func MoveZeroes(nums []int) {
	n := len(nums)
	left, right := 0, 1
	for right < n {
		if nums[left] != 0 {
			left++
		}

		if nums[left] == 0 && nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		right++
	}
}
