package pkg

func QuickSort(nums []int, low, high int) {
	if low < high {
		pivot := partition(nums, low, high)
		QuickSort(nums, low, pivot-1)
		QuickSort(nums, pivot+1, high)
	}

}

func partition(nums []int, low, high int) int {
	pivot := nums[low]

	for low < high {
		for low < high && pivot <= nums[high] {
			high--
		}
		nums[low] = nums[high]

		for low < high && pivot >= nums[low] {
			low++
		}

		nums[high] = nums[low]
	}

	nums[low] = pivot

	return low
}
