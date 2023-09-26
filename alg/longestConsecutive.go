package alg

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
func LongestConsecutive(nums []int) int {
	res := 0
	numSet := make(map[int]bool)

	for _, num := range nums {
		numSet[num] = true
	}

	for num := range numSet {
		if !numSet[num-1] {
			curMax := 1
			for curr := num + 1; numSet[curr]; curr++ {
				curMax++
			}

			res = max(res, curMax)
		}
	}

	return res
}
