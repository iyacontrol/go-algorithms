package alg

func FindMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (right-left)/2 + left

		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}

	}

	return nums[left]
}
