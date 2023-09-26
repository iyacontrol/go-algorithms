package alg

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
func SubarraySum(nums []int, k int) int {
	n := len(nums)

	res := 0
	prefixSum := 0
	hashTable := make(map[int]int)
	hashTable[0] = 1

	for i := 0; i < n; i++ {
		prefixSum = prefixSum + nums[i]
		if v, ok := hashTable[prefixSum-k]; ok {
			res = res + v
		}

		hashTable[prefixSum]++
	}

	return res

}
