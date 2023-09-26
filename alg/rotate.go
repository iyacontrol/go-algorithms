package alg

func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	n := len(nums)
	left, right := 0, n-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
