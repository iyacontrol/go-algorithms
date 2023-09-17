package pkg

import "math"

// 给定一个含有 n 个正整数的数组和一个正整数 target 。

// 找出该数组中满足其总和大于等于 target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

// 示例 1：

// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 示例 2：

// 输入：target = 4, nums = [1,4,4]
// 输出：1

func MinSubArrayLen(target int, nums []int) int {

	n := len(nums)
	res := math.MaxInt

	// 暴力解法
	// sum := 0

	// for i := 0; i < n; i++ {
	// 	for j := i; j < n; j++ {

	// 		for k := i; k <= j; k++ {
	// 			sum += nums[k]
	// 		}

	// 		if sum >= target {
	// 			res = min(res, j-i+1)
	// 		}
	// 		sum = 0
	// 	}
	// }

	// 暴力解法改进
	// sum := 0

	// for i := 0; i < n; i++ {
	// J:
	// 	for j := i; j < n; j++ {

	// 		sum += nums[j]

	// 		if sum >= target {
	// 			res = min(res, j-i+1)
	// 			sum = 0
	// 			// 由于找长度最小的连续子数组，所以一旦找到以后，没有必要继续循环下去
	// 			break J
	// 		}
	// 	}
	// }

	// 前缀和
	// prefixSum := make([]int, n)
	// prefixSum[0] = 0

	// for i := 1; i < n; i++ {
	// 	prefixSum[i] = prefixSum[i-1] + nums[i-1]
	// }

	// for i := 0; i < n; i++ {
	// J:
	// 	for j := i; j >= 0; j-- {
	// 		if nums[i]+prefixSum[i]-prefixSum[j] >= target {
	// 			res = min(res, i-j+1)
	// 			break J
	// 		}
	// 	}
	// }

	// 滑动窗口
	left, right := 0, 0
	sum := 0
	for right < n {
		sum += nums[right]

		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}

		right++
	}

	if res == math.MaxInt {
		return 0
	}
	return res
}
